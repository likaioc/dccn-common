package ankr_const

// RabbitQ name
const TaskManagerQueueName = "task_manager"

// Task Events
const NewTaskEvent = "New"
const UpdateTaskEvent = "Update"
const CancelTaskEvent = "Cancel"

// Hub task status
const TaskStatusNew = "new"
const TaskStatusRunning = "running"
const TaskStatusStartFailed = "startFailed"
const TaskStatusUpdating = "updating"
const TaskStatusUpdateFailed = "updateFailed"
const TaskStatusCancelled = "cancelled"
const TaskStatusDone = "done"

// Data center task status
const DataCenterTaskStartSuccess = "StartSuccess"
const DataCenterTaskStartFailure = "StartFailure"
const DataCenterTaskUpdateSuccess = "UpdateSuccess"
const DataCenterTaskUpdateFailure = "UpdateFailure"
const DataCenterTaskDone = "Done"
const DataCenterTaskCancelled = "Cancelled"

// Data center status
const DataCenteStatusOnLine = "available"
const DataCenteStatusOffLine = "unavailable"

// CLI reply status
const CliReplyStatusSuccess = "Success"
const CliReplyStatusFailure = "Failure"

// CLI errors
const CliErrorReasonDataCenterNotExist = "DataCenter does not exist"
const CliErrorReasonUserNotExist = "User does not exist"
const CliErrorReasonTaskNotExist = "Task does not exist"
const CliErrorReasonUserNotOwn = "User does not own this task"
const CliErrorReasonUpdateFailed = "Task can not be updated"

// To do: Remove this line to generate data center name automatically
const DataCenterName = "datacenter_1"

// To do: Remove this line when usrmgr is ready
const DefaultUserToken = "ed1605e17374bde6c68864d072c9f5c9"

const TaskHidden = "hidden"
const DefaultPort = 50051      // Default port for gRPC connection
const ClientTimeOut = 30       // Default timeout for client connection
const MicroServiceTimeOut = 10 // Default timeout for internal micro service connection
