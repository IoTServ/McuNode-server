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
  local p
  local pin = req.getParam("pin")
  for i=0 , 12 do
	gpio.mode(i, gpio.OUTPUT)
  end
  if (pin~=nil) then
	  res.body="ok"
	  a,b=string.find(pin,'ON')
	  if (a~=nil) then
	  	  c,d=string.find(pin,'%d%d')
		  if (c~=nil) then
			p=string.sub(pin,c,d)
		  else
			p=string.sub(pin,string.find(pin,'%d'))
		  end
		gpio.write(p,1)  
	  else
		c,d=string.find(pin,'%d%d')
		  if (c~=nil) then
			p=string.sub(pin,c,d)
		  else
			p=string.sub(pin,string.find(pin,'%d'))
		  end
		gpio.write(p,0)  
	  end
  elseif (pin==nil) then
	res.file = "indextpl.html"
  end
  return res
end)
mcunode.connect("4567","eiot.club","wifi","123456abc")
