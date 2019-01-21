package ankr_default

import "errors"

// List execution errors
var (
	ErrDataCenterNotExist = errors.New("DataCenter does not exist")
	ErrUserNotExist       = errors.New("User does not exist")
	ErrTaskNotExist       = errors.New("Task does not exist")
	ErrUserNotOwn         = errors.New("User does not own this task")
	ErrUpdateFailed       = errors.New("Task can not be updated")
)
