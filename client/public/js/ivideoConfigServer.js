var serverConfigDialogs = null;

$(function() {
	serverConfigDialogs = new Map();
	serverConfigDialogs.put(SERVER_TYPE_APS, new ViewDialog("apsServerConfigDialog", "btnAPSSaveServer", showConfigAPSServerDlgInfo, {
		url: "/server/saveconfig",
		checkSaveFunc: checkAPSSaveServer,
		getData: getAPSServerPostData,
		getNotifyInfo: getAPSServerNotifyInfo,
		saveSuccessFunc: null
	}));
	enableTabSwitch('apsServerConfigTab');
	serverConfigDialogs.put(SERVER_TYPE_SIP, new ViewDialog("sipServerConfigDialog", "btnSIPSaveServer", showConfigSIPServerDlgInfo, {
		url: "/server/saveconfig",
		checkSaveFunc: checkSIPSaveServer,
		getData: getSIPServerPostData,
		getNotifyInfo: getSIPServerNotifyInfo,
		saveSuccessFunc: null
	}));
	enableTabSwitch('sipServerConfigTab');
		
	serverConfigDialogs.put(SERVER_TYPE_CMS, new ViewDialog("cmsServerConfigDialog", "btnCMSSaveServer", showConfigCMSServerDlgInfo, {
		url: "/server/saveconfig",
		checkSaveFunc: checkCMSSaveServer,
		getData: getCMSServerPostData,
		getNotifyInfo: getCMSServerNotifyInfo,
		saveSuccessFunc: null
	}));
	enableTabSwitch('cmsServerConfigTab');		
});

function showConfigServerDlg(data) {
	var serverConfigInfo = {
		serverInfo: data[0],
		configServerInfo: data[1]
	};
	serverConfigDialogs.get(data[0].Type).show(serverConfigInfo);
}

function getInputLabelInfo(dialogName, id) {
	return $('#' + dialogName + ' label[for="' + id + '"]').html();
}

function checkSaveServerConfig(serverName) {
	var result = true;
	var tabName = serverName + 'ServerConfigTab';
	var dialogName= serverName + 'ServerConfigDialog';
	$('#' + dialogName + ' div input').each(function() {
		var id = $(this).attr('id');
		var info = getInputLabelInfo(dialogName, id);
		var tabId = $(this).parent().parent().parent().attr('id');
		if (!checkInputTextEmptyAndTab(id, info, tabName, tabId)) {
			result = false;
			return false;
		}
	});
	if (!result)
		return false;
	$('#' + dialogName + ' div input[input-int-min]').each(function() {
		var id = $(this).attr('id');
		var info = getInputLabelInfo(dialogName, id);
		var tabId = $(this).parent().parent().parent().attr('id');
		if (!checkInputTextIntegerRangeAndTab(id, info, $(this).attr('input-int-min'), null, tabName, tabId)) {
			result = false;
			return false;
		}
	});
	if (!result)
		return false;
	$('#' + dialogName + ' div input[input-int-max]').each(function() {
		var id = $(this).attr('id');
		var info = getInputLabelInfo(dialogName, id);
		var tabId = $(this).parent().parent().parent().attr('id');
		if (!checkInputTextIntegerRangeAndTab(id, info, null, $(this).attr('input-int-max'), tabName, tabId)) {
			result = false;
			return false;
		}
	});
	if (!result)
		return false;
	$('#' + dialogName + ' div input[id$="Password1"]').each(function() {
		var id = $(this).attr('id');
		var pid = id.substring(0, id.length - 1);
		var info = getInputLabelInfo(dialogName, pid);
		var tabId = $(this).parent().parent().parent().attr('id');
		if (!checkInputTextIdentifyAndTab(pid, id, info, tabName, tabId)) {
			result = false;
			return false;
		}
	});	
	return result;
}

function showConfigAPSServerDlgInfo(serverConfigInfo) {
	$("#apsServerName").html(serverConfigInfo.serverInfo.ServerName);
	$("#apsServerId").val(serverConfigInfo.configServerInfo.APSInfo.Id);
	$("#apsPassword").val(serverConfigInfo.configServerInfo.APSInfo.Password);
	$("#apsPassword1").val(serverConfigInfo.configServerInfo.APSInfo.Password);
	$("#apsRegisterInterval").val(serverConfigInfo.configServerInfo.APSInfo.RegisterInterval);
	$("#apsHeartInterval").val(serverConfigInfo.configServerInfo.APSInfo.DefaultHeartInterval);
	$("#apsAddress").val(serverConfigInfo.configServerInfo.APSInfo.AddressInfo.IP);
	$("#apsPort").val(serverConfigInfo.configServerInfo.APSInfo.AddressInfo.Port);
	$("#apsLogSetupFile").val(serverConfigInfo.configServerInfo.APSInfo.LogInfo.Logcfgfile);
	$("#apsLogFolder").val(serverConfigInfo.configServerInfo.APSInfo.LogInfo.SaveCatalog);
	$("#apsSipServerId").val(serverConfigInfo.configServerInfo.SipServerInfo.Id);
	$("#apsSipServerDomain").val(serverConfigInfo.configServerInfo.SipServerInfo.Domain);
	$("#apsSipServerAddress").val(serverConfigInfo.configServerInfo.SipServerInfo.AddressInfo.IP);
	$("#apsSipServerPort").val(serverConfigInfo.configServerInfo.SipServerInfo.AddressInfo.Port);
	$("#apsCmsServerAddress").val(serverConfigInfo.configServerInfo.CMSServerInfo.Address);
	switchFirstTabSwitch('apsServerConfigTab');
	return true;
}



function checkAPSSaveServer() {
	return checkSaveServerConfig('aps');
}

function getAPSServerPostData() {
	return {
		servername: $('#apsServerName').html(),
		serverId: $("#apsServerId").val(),
		servertype: SERVER_TYPE_APS,
		password: $("#apsPassword").val(),
		registerInterval: $("#apsRegisterInterval").val(),
		defaultHeartInterval: $("#apsHeartInterval").val(),
		address: $("#apsAddress").val(),
		port: $("#apsPort").val(),		
		logcfgfile: $("#apsLogSetupFile").val(),
		saveCatalog: $("#apsLogFolder").val(),
		sipServerId: $("#apsSipServerId").val(),
		sipServerDomain: $("#apsSipServerDomain").val(),		
		sipServerAddress: $("#apsSipServerAddress").val(),
		sipServerPort: $("#apsSipServerPort").val(),	
		cmsServerAddress: $("#apsCmsServerAddress").val(),		
	};
}

function getAPSServerNotifyInfo() {
	return $("#apsServerName").html() + '配置信息';
}



function showConfigSIPServerDlgInfo(serverConfigInfo) {
	$("#sipServerName").html(serverConfigInfo.serverInfo.ServerName);
	$("#sipServerId").val(serverConfigInfo.configServerInfo.SIPInfo.Id);
	$("#sipServerDomain").val(serverConfigInfo.configServerInfo.SIPInfo.Domain);	
	$("#sipPassword").val(serverConfigInfo.configServerInfo.SIPInfo.Password);
	$("#sipPassword1").val(serverConfigInfo.configServerInfo.SIPInfo.Password);
	$("#sipHeartInterval").val(serverConfigInfo.configServerInfo.SIPInfo.DefaultHeartInterval);
	$("#sipAddress").val(serverConfigInfo.configServerInfo.SIPInfo.AddressInfo.IP);
	$("#sipPort").val(serverConfigInfo.configServerInfo.SIPInfo.AddressInfo.Port);
	$("#sipCmsServerAddress").val(serverConfigInfo.configServerInfo.CMSServerInfo.Address);
	$("#sipDbServerAddress").val(serverConfigInfo.configServerInfo.DbInfo.AddressInfo.IP);
	$("#sipDbServerPort").val(serverConfigInfo.configServerInfo.DbInfo.AddressInfo.Port);
	$("#sipDbServerPassword").val(serverConfigInfo.configServerInfo.DbInfo.Password);
	$("#sipDbServerPassword1").val(serverConfigInfo.configServerInfo.DbInfo.Password);
	switchFirstTabSwitch('sipServerConfigTab');
	return true;
}

function checkSIPSaveServer() {
	return checkSaveServerConfig('sip');
}

function getSIPServerPostData() {
	return {
		servername: $('#sipServerName').html(),
		serverId: $("#sipServerId").val(),
		servertype: SERVER_TYPE_SIP,
		domain: $("#sipServerDomain").val(),		
		password: $("#sipPassword").val(),
		defaultHeartInterval: $("#sipHeartInterval").val(),
		address: $("#sipAddress").val(),
		port: $("#sipPort").val(),		
		cmsServerAddress: $("#sipCmsServerAddress").val(),		
		dbServerAddress: $("#sipDbServerAddress").val(),
		dbServerPort: $("#sipDbServerPort").val(),	
		dbServerPassword: $("#sipDbServerPassword").val(),	
	};
}

function getSIPServerNotifyInfo() {
	return $("#sipServerName").html() + '配置信息';
}

function showConfigCMSServerDlgInfo(serverConfigInfo) {
	$("#cmsServerName").html(serverConfigInfo.serverInfo.ServerName);
	$("#cmsDbServerAddress").val(serverConfigInfo.configServerInfo.dbAddress);
	$("#cmsDbServerSchema").val(serverConfigInfo.configServerInfo.dbDefaultSchema);	
	$("#cmsDbServerUser").val(serverConfigInfo.configServerInfo.dbUserName);	
	$("#cmsDbServerPassword").val(serverConfigInfo.configServerInfo.dbPassword);
	$("#cmsDbServerPassword1").val(serverConfigInfo.configServerInfo.dbPassword);
	$("#cmsActiveMQAddress").val(serverConfigInfo.configServerInfo.activemqAddress);	
	switchFirstTabSwitch('cmsServerConfigTab');
	return true;
}

function checkCMSSaveServer() {
	return checkSaveServerConfig('cms');
}

function getCMSServerPostData() {
	return {
		servername: $('#cmsServerName').html(),
		servertype: SERVER_TYPE_CMS,
		dbAddress: $("#cmsDbServerAddress").val(),		
		dbDefaultSchema: $("#cmsDbServerSchema").val(),
		dbUserName: $("#cmsDbServerUser").val(),
		dbPassword: $("#cmsDbServerPassword").val(),
		activemqAddress: $("#cmsActiveMQAddress").val(),
	};
}

function getCMSServerNotifyInfo() {
	return $("#cmsServerName").html() + '配置信息';
}