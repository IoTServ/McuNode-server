local mcunode = require "mcunode"

--url:http://mcunodeserver-ip/proxy/<your-id>/gpio?pin=1&val=0
--return:status=success
mcunode.handle("/gpio",function(req,res)
	local pin = req.getParam("pin")
	local val = req.getParam("val")
	if( pin ~= nil) and (val ~= nil) then
		print("pin: "..pin .." , val: " ..val)
		gpio.mode(pin, gpio.OUTPUT)
		gpio.write(pin, val)
		res.setAttribute("status","success")
	end
	return res
end)
mcunode.handle("/test",function(req,res)
	local name = req.getParam("name")
	res.setAttribute("name",name)
	res.body="wodename:{{name}}"
	return res
end)
--url:http://mcunodeserver-ip/proxy/<your-id>/index.html?name=farry&work=student
mcunode.handle("/index.html",function(req,res)
  res.file = "indextpl.html"
  local pin = req.getParam("pin")
  res.setAttribute("raw",pin)
  gpio.mode(0, gpio.OUTPUT)
  gpio.mode(2, gpio.OUTPUT)
  if( pin == "ON1") then
	gpio.write(0,0)
  elseif(pin == "ON2" ) then
	gpio.write(2,0)
  elseif(pin == "OFF1" ) then
	gpio.write(0,1)
  elseif(pin == "OFF2" ) then
	gpio.write(2,1)
  end
  return res
end)
mcunode.connect("4567","eiot.club","wifi","123456abc")