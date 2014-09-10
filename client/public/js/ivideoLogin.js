$(function() {
	$("form").submit = function() {
		if (!checkInputTextEmpty("username", "用户名")) {
			return false;
		}
		if (!checkInputTextEmpty("password", "密码")) {
			return false;
		}
		return true;
	};
});