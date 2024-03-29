package advancedlog

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type CloseBuilder struct {
	once    sync.Once
	toClose []Closer
}

func (c *CloseBuilder) GetCloser() Closer {
	return func(e *logrus.Entry) error {
		c.once.Do(func() {
			// Clause toCLose slice, so we'll be able to add new Closers
			for _, closer := range c.toClose {
				err := closer(e)
				if err != nil {
					e.Errorf("closing: %s", err)
				}
			}
		})
		return nil
	}
}

func (c *CloseBuilder) AddClose(closer func() error, serviceName string) {
	c.toClose = append(c.toClose, func(entry *logrus.Entry) error {
		entry.Infof("Closing %s!", serviceName)
		err := closer()
		if err != nil {
			entry.Errorf("closing %s: %s", serviceName, err)
			// No need to return err to exec function
			return nil
		}
		entry.Infof("Closed %s!", serviceName)
		return nil
	})
}

func (c *CloseBuilder) AddCloser(closer Closer) {
	c.toClose = append(c.toClose, func(entry *logrus.Entry) error {
		err := closer(entry)
		if err != nil {
			return err
		}
		return nil
	})
}
