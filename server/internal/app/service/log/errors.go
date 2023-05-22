package log

import "k8s.io/klog"

func ErrorLog(err error) error {
	if err != nil {
		klog.Error(err)
		return err
	}
	return nil
}
func ErrStrLog(err string) string {
	if err != "" {
		klog.Error(err)
		return err
	}
	return ""
}
