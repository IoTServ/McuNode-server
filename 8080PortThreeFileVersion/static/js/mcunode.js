var websocket;
window.addEventListener('load', function() {
        var wsUri ="ws://"+window.location.host+"/ws/"+id;
        var term = new Terminal({cols: 100,rows: 35});
        term.open(document.getElementById('terminal'));
        function runWebSocket() {
            websocket = new WebSocket(wsUri);
            websocket.onopen = function(evt) {
                onOpen(evt)
            };
            websocket.onclose = function(evt) {
                onClose(evt)
            };
            websocket.onmessage = function(evt) {
                onMessage(evt)
            };
            websocket.onerror = function(evt) {
                onError(evt)
            };
        }

        function onOpen(evt) {
            writeToScreen("CONNECTED");
            doSend("=wifi.sta.getip()");
        }

        function onClose(evt) {
            writeToScreen("DISCONNECTED");
        }

        function onMessage(evt) {
            s=String(evt.data);
            console.log('row data:'+s);
            //s=s.substring(0,s.length-3);
            s=s.replace(/\n/g,'\r\n');
            writeToScreen('>: '+ s);
            //websocket.close();
        }

        function onError(evt) {
            writeToScreen('ERROR: '+ evt.data);
        }

        function doSend(message) {
            writeToScreen("SENT: " + message);
            websocket.send(message);
        }

        function writeToScreen(message) {
            term.write(message+'\r\n>');
        }
        runWebSocket();
		//

        var send_data='';
        term.on('data', function(data) {
            if (data!='\r'&&data!='\n') {
                send_data = send_data + data;
                term.write(data);
                console.log(data);
            }
            if (data=='\r'||data=='\n'){
                term.write('\r\n>');
                console.log(send_data);

                var strs=send_data.split(" "); //字符分割
                switch(strs[0]){
                    case 'ls':
                        send_data = 'ls()';
                        break;
                    case 'reboot':
                        send_data = 'node.restart()';
                        break;
                    case 'rm':
                        var fname=strs[1];
                        send_data = 'file.remove(\''+fname+'\')';
                        break;
                    case 'cat':
                        var fname=strs[1];
                        send_data = 'cat(\''+fname+'\')';
                        break;
                    case 'dl':
                        var host;
                        var uri;
                        var fname;
                        var url=strs[1];
                        urlstrs=send_data.split("/");
                        host=urlstrs[2];
                        uri=urlstrs[3]+'/'+urlstrs[4];
                        fname=urlstrs[4];
                        if (typeof(strs[2])!='undefined'){
                            fname=strs[2];
                        }
                        //from host+uri download a file,save as fname
                        send_data = 'downloadfile(\''+host+'\',\''+uri+'\',\''+fname+'\')';
                        break;
                    case 'run':
                        var host;
                        var uri;
                        var fname;
                        var url=strs[1];
                        urlstrs= new Array();
                        urlstrs=send_data.split("/");
                        host=urlstrs[2];
                        uri=urlstrs[3]+'/'+urlstrs[4];
                        fname=urlstrs[4];
                        //from host+uri download a file,save as fname
                        message = 'downloadfile(\''+host+'\',\''+uri+'\',\''+'run.lua\')';
                        websocket.send(message);
                        send_data = 'dofile(\'run.lua\')';
                        break;
                    case 'dofile':
                        var fname=strs[1];
                        send_data = 'dofile(\''+fname+'\')';
                        break;
                    default:
                        console.log('没有可以匹配的命令'+send_data);
                    }
                websocket.send(send_data);
                send_data='';
            }
        });


        term.write('Welcome to use \033[1;3;31m McuNode!\033[0m \r\n>');
    }, false);



function writeToMcu(){
            //var code=document.getElementById("code").value.split('\n');
            var code  = editor.getValue().split('\n');
            lineleng=code.length;
            console.log(lineleng);
            filename=document.getElementById("filename").value||'fogot.lua';
            websocket.send('file.remove(\''+filename+'\', \'w+\')');
            websocket.send('file.open(\''+filename+'\', \'w+\')');
            var innerI = 0;
            mytime = setInterval(function () { write() }, 200);
            function write() {
                if(innerI<lineleng){
                code[innerI]=code[innerI].replace(/\'/g,'\"').replace(/\"/g,'\\'+'\"');
                console.log('file.writeline(\"'+code[innerI]+'\")');
                websocket.send('file.writeline(\"'+code[innerI]+'\")');}
                if(innerI == lineleng){
                    websocket.send('file.close()');console.log('closed');
                }
                if (innerI == lineleng+1) {clearInterval(mytime);}
                innerI++;
            }

        }
