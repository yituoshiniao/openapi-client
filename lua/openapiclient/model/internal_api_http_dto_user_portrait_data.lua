--[[
  goods-center-logic API

  goods-center-logic服务文档

  The version of the OpenAPI document: 1.0
  
  Generated by: https://openapi-generator.tech
]]

-- internal_api_http_dto_user_portrait_data class
local internal_api_http_dto_user_portrait_data = {}
local internal_api_http_dto_user_portrait_data_mt = {
	__name = "internal_api_http_dto_user_portrait_data";
	__index = internal_api_http_dto_user_portrait_data;
}

local function cast_internal_api_http_dto_user_portrait_data(t)
	return setmetatable(t, internal_api_http_dto_user_portrait_data_mt)
end

local function new_internal_api_http_dto_user_portrait_data(country, last_login, user_id, vip_info)
	return cast_internal_api_http_dto_user_portrait_data({
		["country"] = country;
		["last_login"] = last_login;
		["user_id"] = user_id;
		["vip_info"] = vip_info;
	})
end

return {
	cast = cast_internal_api_http_dto_user_portrait_data;
	new = new_internal_api_http_dto_user_portrait_data;
}
