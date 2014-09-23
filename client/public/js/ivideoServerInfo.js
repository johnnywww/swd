function getServerInfoAjaxErrorFunc(item, successFunc, errorFunc) {
	$.get('/server/info/' + item, {
			servername: $("#servernameId").text()
		},
		//回调函数
		function(data) {
			if (!data.Succ) {
				if (errorFunc) {
					errorFunc(data);
				}
			} else {
				successFunc(data);
			}
		},
		"json"
	);
}

function getServerInfoAjax(item, successFunc, showError) {
	var errorFunc = null;
	if (showError) {
		errorFunc = function(data) {
			notifyErrorMsg(data.Msg);
		};
	}
	getServerInfoAjaxErrorFunc(item, successFunc, errorFunc);
}

function getProcessInfo(showError) {
	var errorFunc = null;
	if (showError) {
		errorFunc = function(data) {
			notifyErrorMsg(data.Msg);
		};
	} else {
		errorFunc = function() {
			$("#process-cpu").text("");
			$("#process-mem").text("");
			$("#process-realmem").text("");
			$("#process-virtualmem").text("");
			$("#process-state").text("");
		}
	}
	getServerInfoAjaxErrorFunc('process', function(data) {
		$("#process-cpu").text(data.Data.cpuRate);
		$("#process-mem").text(data.Data.memRate);
		$("#process-realmem").text(data.Data.realMem);
		$("#process-virtualmem").text(data.Data.virtualMem);
		$("#process-state").text(data.Data.state);
	}, errorFunc);
}

function getGeneralInfo(showError) {
	getServerInfoAjax('general', function(data) {
		$("#os-hostname").text(data.Data.hostname);
		$("#os-uptime").text(data.Data.runtime);
		$("#os-time").text(data.Data.time);
		$("#os-kernel").text(data.Data.kernel);
	}, showError);
}

function getCpuInfo(showError) {
	getServerInfoAjax('cpu', function(data) {
		$("#cpu-idle").text(data.Data.idle);
		$("#cpu-busy").text(data.Data.busy);
		$("#cpu-user").text(data.Data.user);
		$("#cpu-iowait").text(data.Data.iowait);
		$("#cpu-system").text(data.Data.system);
		$("#cpu-nice").text(data.Data.nice);
		$("#cpu-irq").text(data.Data.irq);
		$("#cpu-softirq").text(data.Data.softirq);
		$("#cpu-steal").text(data.Data.steal);
		$("#cpu-guest").text(data.Data.guest);
	}, showError);
}

function clearStatusInfos() {
	$("#serverStatusTab").html('');
	$("#serverStatusTabContent").html('');
}

function addStatusInfoTab(key) {
	var str = '<li><a href="#status-' + key + '">' + key + '</a></li>';
	$("#serverStatusTab").append(str);
}

function addStatusInfoContent(key, value) {
	var str = '<div id="status-' + key + '" class="tab-pane">';
	str += '<table class="table table-bordered"><thead><tr><th>信息</th></tr></thead>';
	str += '<tbody>';
	if (null != value) {
		for (var index in value) {
			str += '<tr><td>' + value[index].Info + '</td></tr>';
		}
	}
	else {
		str += '<tr><td>&nbsp;</td></tr>';
	}
	str += '</tbody>';
	str += '</table>';
	str += '</div>';
	$("#serverStatusTabContent").append(str);
}

function serverStatusTabSwitch() {
	enableTabSwitch('serverStatusTab');
}

function getStatusInfo(showError) {
	getServerInfoAjax('status', function(data) {
		clearStatusInfos();
		var vExist = false;
		for (key in data.Data) {
			addStatusInfoTab(key);
			addStatusInfoContent(key, data.Data[key]);
			vExist = true;
		}
		if (!vExist) {
			addStatusInfoTab('Basic');
			addStatusInfoContent('Basic', null);
		}
		serverStatusTabSwitch();
	}, showError);
}


function getMemInfo(showError) {
	getServerInfoAjax('mem', function(data) {
		var ram_total = data.Data[0];
		var ram_used = Math.round((data.Data[1] / ram_total) * 100);
		var ram_free = Math.round((data.Data[2] / ram_total) * 100);

		$("#mem-total").text(sprintf('%.3f', ram_total / (1024 * 1024)));
		$("#mem-used").text(sprintf('%.3f', data.Data[1] / (1024 * 1024)));
		$("#mem-free").text(sprintf('%.3f', data.Data[2] / (1024 * 1024)));

		$("#mem-freerate").text(ram_free);
		$("#mem-usedrate").text(ram_used);
	}, showError);
}

function addNetCardTr(value) {
	var str = '<tr>';
	for (var i = 0; i < value.length; i++) {
		str += '<td>' + value[i] + '</td>';
	}
	str += '</tr>';
	$("#tbnetcard tbody").append(str);
}

function getNetCardInfo(showError) {
	getServerInfoAjax('netcard', function(data) {
		clearTable("netcard");
		for (var i = 0; i < data.Data.length; i++)
			addNetCardTr(data.Data[i]);
		$(window).resize();
	}, showError);
}

function getAllInfo() {
	getGeneralInfo(false);
	getCpuInfo(false);
	getMemInfo(false);
	getNetCardInfo(false);
	getProcessInfo(false);
	getStatusInfo(false);
}

var funcMap = {
	general: getGeneralInfo,
	cpu: getCpuInfo,
	all: getAllInfo,
	mem: getMemInfo,
	net: getNetCardInfo,
	process: getProcessInfo,
	status: getStatusInfo,
};

function checkBoxObjChecked(obj) {
	return true == obj.prop("checked");
}

function refreshFunc() {
	getAllInfo();
	if (checkBoxObjChecked($("#autoRefreshCheck"))) {
		setTimeout(function() {
			refreshFunc();
		}, $("#refresh-interval").val() * 1000);
	}
}

$(function() {
	bindClickFunc($('div a>i.icon-refresh'), function(obj, event) {
		var id = obj.parent().attr('id');
		var item = id.split("-").splice(-1)[0];
		if ('all' == item) {
			funcMap[item]();
		} else {
			funcMap[item](true);
		}
	});
	$('#autoRefreshCheck').bind("click", function(e) {
		if (checkBoxObjChecked($(this))) {
			$("#refresh-interval").prop("disabled", false);
			refreshFunc();		
		} else {
			$("#refresh-interval").prop("disabled", true);
		}
	});
	refreshFunc();
});