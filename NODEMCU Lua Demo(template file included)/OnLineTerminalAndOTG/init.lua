id = 'yourid'  --Device ID
ssid = 'yourSSID'  --WiFi Name
ssidpwd = 'WiFiPassword'  --view on http://www.mcunode.com/proxy/<yourid>/<anystring> like:http://www.mcunode.com/proxy/4567/index.html
ip="yourServerIpOrHostName"  --Eg："192.168.1.105" or“www.mcunode.com”
function ls()
	local l = file.list()
	for k,v in pairs(l) do
		print("name:"..k..", size:"..v)
	end
end
function cat(filename)
	local line
	file.open(filename, "r")
	while 1 do
		line = file.readline() 
		if line == nil then
			break
		end
		line = string.gsub(line, "\n", "")
		print(line)
	end
	file.close()
end
function startServer()
print(wifi.sta.getip())
sk=net.createConnection(net.TCP, 0)
sk:on("receive", function(sck, c) node.input(c) end )   --print(c)
sk:on("connection", function(sck, c) 
--print(c)
sk:send(id)
tmr.alarm(2, 30000, 1, function() 
	sk:send('<h1></h1>')
end)
function s_output(str)
         if (sk~=nil and str~='')    then
            sk:send(str)
         end
      end
node.output(s_output,1)
end )
sk:on("disconnection",function(conn,c) 
         --node.output(nil) 
		 print('reconnect')
		 sk:connect(8001,ip)
		 sk:send(id)
      end)
sk:connect(8001,ip)
end
wifi.setmode(wifi.STATION)
wifi.sta.config(ssid,ssidpwd)    --set your ap info !!!!!!
wifi.sta.autoconnect(1)
tmr.alarm(1, 1000, 1, function() 
   if wifi.sta.getip()==nil then
      print("Connect AP, Waiting...") 
   else
      startServer()
      tmr.stop(1)
   end
end)