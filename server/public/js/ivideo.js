function indexLogout() {
	$("#logoutForm").submit()
}

function adjustIFrameHeight(id) {
	$("#" + id).load(function() {
		var ifm = document.getElementById(id);
		var subWeb = document.frames ? document.frames[id].document : ifm.contentDocument;
		if (ifm != null && subWeb != null) {
			ifm.height = subWeb.body.scrollHeight;
			$(this).css("height", ifm.height);
		}
	});
}

function loadIFrameSrc(id, src) {
	$("#" + id).attr("src", src);
}

function notifyMsg(msg, type1) {
	noty({
		text: msg,
		layout: "center",
		type: type1,
		timeout: 2000
	});
}

function notifySuccessMsg(msg) {
	notifyMsg(msg, "success");
}

function notifyErrorMsg(msg) {
	notifyMsg(msg, "error");
}

function notifyInfoMsg(msg) {
	notifyMsg(msg, "info");
}


//ajaxify menus
$('a.ajax-link').click(function(e) {
	if ($.browser.msie) e.which = 1;
	if (e.which != 1 || $(this).parent().hasClass('active')) return;
	var $clink = $(this);
	$('ul.main-menu li.active').removeClass('active');
	$clink.parent('li').addClass('active');
});

(function($) {
	$.ajaxDelete = function(url, data, callback, type) {
		// shift arguments if data argument was omitted
		if (jQuery.isFunction(data)) {
			type = type || callback;
			callback = data;
			data = undefined;
		}

		return jQuery.ajax({
			type: "DELETE",
			url: url,
			data: data,
			success: callback,
			dataType: type
		});
	};
})(jQuery);


function Map() {
	this.elements = new Array();

	//获取MAP元素个数 
	this.size = function() {
		return this.elements.length;
	}

	//判断MAP是否为空 
	this.isEmpty = function() {
		return (this.elements.length < 1);
	}

	//删除MAP所有元素 
	this.clear = function() {
		this.elements = new Array();

	}
	//向MAP中增加元素（key, value)  
	this.put = function(_key, _value) {
		//alert(this.containsKey(_key));
		if (this.containsKey(_key)) {
			this.remove(_key);
		}
		this.elements.push({
			key: _key,
			value: _value
		});
	}

	//删除指定KEY的元素，成功返回True，失败返回False 
	this.remove = function(_key) {
		var bln = false;
		try {
			for (i = 0; i < this.elements.length; i++) {
				if (this.elements[i].key == _key) {
					this.elements.splice(i, 1);
					return true;
				}
			}
		} catch (e) {
			bln = false;
		}
		return bln;
	}

	//获取指定KEY的元素值VALUE，失败返回NULL 
	this.get = function(_key) {
		try {
			for (i = 0; i < this.elements.length; i++) {
				if (this.elements[i].key == _key) {
					return this.elements[i].value;
				}
			}
		} catch (e) {
			return null;
		}
	}

	//获取指定索引的元素（使用element.key，element.value获取KEY和VALUE），失败返回NULL 
	this.element = function(_index) {
		if (_index < 0 || _index >= this.elements.length) {
			return null;
		}
		return this.elements[_index];
	}

	//判断MAP中是否含有指定KEY的元素 
	this.containsKey = function(_key) {
		var bln = false;
		try {
			for (i = 0; i < this.elements.length; i++) {
				if (this.elements[i].key == _key) {
					bln = true;
				}
			}
		} catch (e) {
			bln = false;
		}
		return bln;
	}

	//判断MAP中是否含有指定VALUE的元素 
	this.containsValue = function(_value) {
		var bln = false;
		try {
			for (i = 0; i < this.elements.length; i++) {
				if (this.elements[i].value == _value) {
					bln = true;
				}
			}
		} catch (e) {
			bln = false;
		}
		return bln;
	}

	//获取MAP中所有VALUE的数组（ARRAY） 
	this.values = function() {
		var arr = new Array();
		for (i = 0; i < this.elements.length; i++) {
			arr.push(this.elements[i].value);
		}
		return arr;
	}

	//获取MAP中所有KEY的数组（ARRAY） 
	this.keys = function() {
		var arr = new Array();
		for (i = 0; i < this.elements.length; i++) {
			arr.push(this.elements[i].key);
		}
		return arr;
	}

}

function ServerTypeMap() {
	this.datas = new Map();
	this.datas.put(SERVER_TYPE_CMS, '中心管理服务器');
	this.datas.put(SERVER_TYPE_SIP, 'SIP服务器');
	this.datas.put(SERVER_TYPE_MTS, '转发服务器');
	this.datas.put(SERVER_TYPE_APS, '报警服务器');
	this.element = function(index) {
		return this.datas.element(index)
	}
	this.keys = function() {
		return this.datas.keys();
	}
	this.values = function() {
		return this.datas.values();
	}
	this.size = function() {
		return this.datas.size();
	}
	this.get = function(key) {
		return this.datas.get(key);
	}
}

var serverTypeMap = new ServerTypeMap();

function checkValueEmpty(value, info) {
	if ('' == value) {
		notifyErrorMsg('请输入' + info);
		return false;
	}
	return true;
}

function checkInputTextEmpty(id, info) {
	if (!checkValueEmpty($(id).val(), info)) {
		$(id).focus();
		return false;
	}
	return true;
}

function checkValueIdentify(value, value1, info) {
	if (value != value1) {
		notifyErrorMsg(info + '输入不一致');
		return false;
	}
	return true;
}

function isAdminUserName(value) {
	return 'admin' == value;
}

function bindClickFunc(obj, func) {
	obj.unbind("click");
	obj.bind("click", function(e) {
		e.preventDefault();
		func($(this));
	});
}

function dialogSaveClick(e) {
	if (e.data) {
		e.data.save();
	}
}

function getTrByTd0Text(tableName, tdText) {
	var result = null;
	var trName = "#tb" + tableName + " tr";
	$(trName).each(function() {
		if ($(this).children('td').eq(0).text() == tdText) {
			result = $(this);
		}
	});
	return result;
}

function bindTdEdtDelBtnFunc(tableName, delFunc, editFunc) {
	var delBtnId = "#tb" + tableName + " td a.btn-danger";
	var edtBtnId = "#tb" + tableName + " td a.btn-info";
	$(delBtnId).each(function() {
		bindClickFunc($(this), delFunc);
	});

	$(edtBtnId).each(function() {
		bindClickFunc($(this), editFunc);
	});
}

function ViewDialog(dialogId, btnId, showFunc, saveData) {
	this.dialogId = dialogId;
	this.showFunc = showFunc;
	this.saveData = saveData;


	this.show = function(data) {
		if (this.showFunc(data)) {
			$('#' + this.dialogId).modal("show");
		}
	}

	this.save = function() {
		if (!this.saveData)
			return;
		if (!this.saveData.checkSaveFunc()) {
			return;
		}
		confirmDialogFunc('确定保存' + this.saveData.getNotifyInfo() + '吗?', function(e) {
			var viewDialog = e.data;
			$.post(viewDialog.saveData.url, viewDialog.saveData.getData(),
				//回调函数
				function(data) {
					console.log(data)
					hideConfirmDialog();
					if (!data.Succ) {
						notifyErrorMsg(data.Msg);
					} else {
						viewDialog.hide();
						notifySuccessMsg("保存" + viewDialog.saveData.getNotifyInfo() + "成功");
						if (viewDialog.saveData.saveSuccessFunc) {
							viewDialog.saveData.saveSuccessFunc(data.Data);
						}
					}
				},
				//返回类型
				"json"
			);
		}, this);
	}

	$('#' + btnId).bind("click", this, dialogSaveClick);

	this.hide = function() {
		$('#' + this.dialogId).modal("hide");
	}
}