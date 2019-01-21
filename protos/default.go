package ankr_default

// RabbitQ name
const MQDeployTask = "topic.deploy.task"
const MQFeedbackTask = "topic.feedback.task"

// To do: Remove this line when usr mgr is ready
const Secret = "ed1605e17374bde6c68864d072c9f5c9"

const ServiceTaskMgrPort = 50051 // Default port for gRPC connection
const ServiceUserMgrPort = 50052 // Default port for gRPC connection
const ServiceEmailPort = 50053   // Default port for gRPC connection
const ServiceDcMgrPort = 50054   // Default port for gRPC connection

// Registry server name
const (
	TaskMgrRegistryServerName = "go.micro.srv.v1.task"
	UserMgrRegistryServerName = "go.micro.srv.v1.user"
	EmailRegistryServerName   = "go.micro.srv.v1.email"
	DcMgrRegistryServerName   = "go.micro.srv.v1.data.center"
)
