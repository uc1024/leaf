{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

type I{{$svrType}}Module interface{
    {{range .Methods}}
    {{.Comment}}
    {{.Name}}(ctx context.Context,req *{{.Request}}, rsp *{{.Reply}})(err error)
    {{end}}
}

{{/* 注册leaf 消息*/}}
// * 消息注册
func Register{{$svrType}}Message(processor *iprotobuf.Processor) {
   {{- range .Methods}}
    {{/* 过滤code是0的消息注册 */}}
    {{- if ne .RequestCode 0}}
    processor.Register({{.RequestCode}},&{{.Request}}{})
    {{- end }}
    {{- if ne .ReplyCode 0}}
    processor.Register({{.ReplyCode}},&{{.Reply}}{})
    {{- end }}
   {{- end }}
}

{{/* 注册leaf 服务响应模式*/}}
// * RegisterLeaf 响应消息初始化
func Register{{$svrType}}ModuleMessage(processor *iprotobuf.Processor, skeleton *module.Skeleton,m I{{$svrType}}Module,rpc *chanrpc.Server) {
   Register{{$svrType}}Message(processor)
   {{- range .Methods}}
	processor.SetRouter(&{{.Request}}{}, rpc)
    skeleton.RegisterChanRPC(reflect.TypeOf(&{{.Request}}{}), func(args []interface{}){
	request := args[0].(*{{.Request}})
    _ = request
	// * agent conn
	agent := args[1].(gate.Agent) 
    {{if ne .NotAuth true }}
{{/*  未登录断开连接*/}}
    if !agent.Auth() {
        agent.Close()
        return
    } 
    {{- end}}
	ctx := context.WithValue(context.Background(), "agent", agent)
    response:=&{{.Reply}}{}
{{/* 调用对应函数 */}}
    err := m.{{.Name}}(ctx,request,response)
    if err != nil {
        if agent.Auth(){
           agent.WriteMsg(response)
           return
        }
        agent.Close()
        return
    }
    agent.WriteMsg(response)
    })
   {{- end}}
}