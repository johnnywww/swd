function addServerTypeList() {
	var $ul = $("#servertype");
	for (i = 0; i < serverTypeMap.size(); i++) {
		var $htmlLi = $('<option value="' + serverTypeMap.element(i).key + '">' + serverTypeMap.element(i).value + '</option>'); //创建DOM对象
		$ul.append($htmlLi);
	};
}

function showServerTypeBtnValue(type1) {
	$('#servertype').val(type1);
}

function showServerDialog(serverInfo) {
	if (serverInfo) {
		$('#servername').val(serverInfo.serverName);
		$('#oservername').val(serverInfo.serverName);
		$('#address').val(serverInfo.address);
		showServerTypeBtnValue(serverInfo.type);
	} else {
		$('#servername').val("");
		$('#oservername').val("");
		$('#address').val("");
		showServerTypeBtnValue(0);
	}
	// $('#serverDialog div.modal-footer a.btn-primary').click(sureFunc);
	return true;
}

function checkSaveServer() {
	if (!checkInputTextEmpty("servername", "名称")) {
		return false;
	}
	if (!checkInputTextEmpty("address", "地址")) {
		return false;
	}
	if (!checkValueEmpty($('#servertype').val(), "服务器类型")) {
		return false;
	}
	return true;
}


function getServerPostData() {
	return {
		servername: $('#servername').val(),
		oservername: $('#oservername').val(),
		address: $("#address").val(),
		servertype: $('#servertype').val()
	};
}

function getServerNotifyInfo() {
	return $("#servername").val();
}

var serverDialog = null;

$(function() {
	addServerTypeList();
	serverDialog = new ViewDialog("serverDialog", "btnSaveServer", showServerDialog, {
		url: "/server/save",
		checkSaveFunc: checkSaveServer,
		getData: getServerPostData,
		getNotifyInfo: getServerNotifyInfo,
		saveSuccessFunc: null
	});
});

function getServerDialog() {
	return serverDialog;
}

function bindServerTableTdBtnFunc() {
	bindTdEdtDelBtnFunc("Servers", deleteServer, showEditServerDlg);
	var configBtnId = "#tbServers td a.btn-primary";

	$(configBtnId).each(function() {
		bindClickFunc($(this), getConfigServer);
	});
}

function addServerTb(data) {
	$str = '<tr>'
	$str += '<td class="sorting_1">' + data.ServerName + '</td>';
	$str += '<td class="center ">' + serverTypeMap.get(data.Type) + '</td>';
	$str += '<td class="center" servername-value="' + data.ServerName + '">';
	$str += '<a href="#" class="btn btn-primary"><i class="icon-cog icon-white"></i>配置</a>';
	$str += '<a href="#" servertype-value="' + data.Type + '" address-value="' + data.Address + '" action-value="e" class="btn btn-info"><i class="icon-edit icon-white"></i>\n编辑</a>\n';
	$str += '<a href="#" action-value="d" class="btn btn-danger"><i class="icon-trash icon-white"></i>\n删除</a>';
	$str += '</td>';
	$str += '</tr>';
	$("#tbServers").append($str);
	bindServerTableTdBtnFunc();
}

function getConfigServer(hrefObj, event) {
	var serverName = hrefObj.parent().attr('servername-value');
	$.get('/serverconfig', {
			servername: serverName
		},
		//回调函数
		function(data) {
			if (!data.Succ) {
				notifyErrorMsg(data.Msg);
			} else {
				showConfigServerDlg(data.Data);
			}
		},
		//返回类型
		"json"
	);
}


function showEditServerDlg(hrefObj, event) {
	var serverInfo = {
		serverName: hrefObj.parent().attr('servername-value'),
		address: hrefObj.attr('address-value'),
		type: hrefObj.attr('servertype-value')
	}
	getServerDialog().show(serverInfo);
}


function changeServerProfile(data) {
	if (data.AddFlag) {
		addServerTb(data);
	} else {
		var tr = getTrByTd0Text("Servers", data.OServerName);
		if (tr) {
			tr.children('td').eq(0).text(data.ServerName);
			tr.children('td').eq(1).text(serverTypeMap.get(data.Type));
			tr.children('td').eq(2).attr('servername-value', data.ServerName);
			tr.children('td').eq(2).find('a.btn-info').attr('servertype-value', data.Type);
			tr.children('td').eq(2).find('a.btn-info').attr('address-value', data.Address);
		}
	}
}

function deleteServer(hrefObj, event) {
	serverName = hrefObj.parent().attr('servername-value');
	confirmDialogFunc('确定删除' + serverName + '吗?', function(e) {
		//向add.ashx请求结果
		$.post('/server/delete', {
				servername: serverName
			},
			//回调函数
			function(data) {
				console.log(data)
				hideConfirmDialog();
				if (!data.Succ) {
					notifyErrorMsg(data.Msg);
				} else {
					hrefObj.parent().parent().remove();
					notifySuccessMsg("删除服务器成功");
				}
			},
			//返回类型
			"json"
		);
	});
}