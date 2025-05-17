package main

import "fmt"

// 定义工作流程的接口
type Workflow interface {
	Start()
	Execute()
	End()
}

// 模板方法：作为独立函数实现运行流程
func RunWorkflow(w Workflow) {
	w.Start()
	w.Execute()
	w.End()
}

// 第一個具体的工作流程
type ReportWorkflow struct{}

func (r *ReportWorkflow) Start() {
	fmt.Println("工作流程开始")
}

func (r *ReportWorkflow) Execute() {
	fmt.Println("执行报告生成流程")
}

func (r *ReportWorkflow) End() {
	fmt.Println("工作流程结束")
}

// 第二个具体的工作流程
type AuditWorkflow struct{}

func (a *AuditWorkflow) Start() {
	fmt.Println("工作流程开始")
}

func (a *AuditWorkflow) Execute() {
	fmt.Println("执行审核流程")
}

func (a *AuditWorkflow) End() {
	fmt.Println("工作流程结束")
}

func main() {
	report := &ReportWorkflow{}
	RunWorkflow(report)

	audit := &AuditWorkflow{}
	RunWorkflow(audit)
}
