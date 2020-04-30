package sshclient

type AgentTasks struct {
	tasks []*AgentTask
}

func (t *AgentTasks) NewTask(cmd AgentCommand) *AgentTask {

	task := NewAgentTask(cmd)

	t.tasks = append(t.tasks, task)

	return task

}

type TransferTask struct {
	Dir  bool
	Src  string
	Dest string
}

type TaskHook struct {
	Upload   []TransferTask
	Download []TransferTask

	ConnectIB    bool
	DisconnectIB bool
}

func (h *TaskHook) upload(src, dest string, dir bool) *TaskHook {

	h.Upload = append(h.Upload, TransferTask{
		Dir:  dir,
		Src:  src,
		Dest: dest,
	})

	return h

}

func (h *TaskHook) download(src, dest string, dir bool) *TaskHook {

	h.Download = append(h.Download, TransferTask{
		Dir:  dir,
		Src:  src,
		Dest: dest,
	})

	return h

}

func (h *TaskHook) UploadFile(src, dest string) *TaskHook {

	return h.upload(src, dest, false)
}

func (h *TaskHook) DownloadFile(src, dest string) *TaskHook {

	return h.download(src, dest, false)
}

func (h *TaskHook) UploadDir(src, dest string) *TaskHook {

	return h.upload(src, dest, true)
}

func (h *TaskHook) DownloadDir(src, dest string) *TaskHook {

	return h.download(src, dest, true)
}

type AgentTask struct {
	cmd        AgentCommand
	OnError    OnErrorRespond
	OnSuccess  OnSuccessRespond
	OnProgress OnProgressRespond

	After  *TaskHook
	Before *TaskHook

	body []byte
	err  *error
}

func NewAgentTask(cmd AgentCommand) *AgentTask {

	task := &AgentTask{
		cmd: cmd,
	}

	var body []byte

	task.OnSuccess, task.OnError, task.OnProgress = successChecker(&body, task.err)
	task.body = body

	return task
}
