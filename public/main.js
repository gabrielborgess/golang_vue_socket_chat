const socket=io();

new Vue({
    el:'#chat-app',
    created(){
        socket.on("chat message",(msg)=>{
            this.mensajes.push({
                text:msg,
                date:new Date().toLocaleDateString()
            })
        })

    },
    data:{
        mensaje:'',
        mensajes:[]
    },
    methods:{
        enviarMensaje(){
            socket.emit('chat message',this.mensaje)
            this.mensaje='';
        }

    }
    })