function showUserDialog(userInfo) {
    var username = "";
    if (userInfo)
        username = userInfo.userName;
    $('#username').val(username);
    if (isAdminUserName(username)) {
        $('#username').attr("disabled", true);
    } else {
        $('#username').attr("disabled", false);
    }

    $('#ousername').val(username);
    $('#password').val("");
    $('#password1').val("");
    return true;
}

function checkSaveUser() {
    if (!checkInputTextEmpty("username", "名称")) {
        return false;
    }
    if (!checkInputTextEmpty("password", "密码")) {
        return false;
    }
    if (!checkInputTextEmpty("password1", "再次输入密码")) {
        return false;
    }
    
    if (!checkInputTextIdentify("password", "password1", '密码')) {
        return false;
    }

    if (isAdminUserName($("#username").val()) && !isAdminUserName($("#ousername").val())) {
        $("#username").focus();
        notifyErrorMsg('不能创建管理员账号');
        return false;
    }
    return true;
}

function getUserPostData() {
    return {
        username: $('#username').val(),
        ousername: $('#ousername').val(),
        password: $("#password").val(),
    };
}

function getUserNotifyInfo() {
    return $("#username").val();
}

var userDialog = null;

$(function() {
    userDialog = new ViewDialog("userDialog", "btnSaveUser", showUserDialog, {
        url: "/user/save",
        checkSaveFunc: checkSaveUser,
        getData: getUserPostData,
        getNotifyInfo: getUserNotifyInfo,
        saveSuccessFunc: null
    });
});

function getUserDialog() {
    return userDialog;
}

function addUserTb(data) {
    $str = '<tr>'
    $str += '<td class="sorting_1">' + data.UserName + '</td>';
    $str += '<td class="center" username-value="' + data.UserName + '">';
    $str += '<a href="#" action-value="e" class="btn btn-info"><i class="icon-edit icon-white"></i>\n编辑</a>\n';
    $str += '<a href="#" action-value="d" class="btn btn-danger"><i class="icon-trash icon-white"></i>\n删除</a>';
    $str += '</td>';
    $str += '</tr>';
    $("#tbUsers").append($str);
    bindTdEdtDelBtnFunc("Users", deleteUser, showEditUserDlg);
}

function changeUserProfile(data) {
    if (data.AddFlag) {
        addUserTb(data);
    } else {
        var tr = getTrByTd0Text("Users", data.OUserName);
        if (tr) {
            tr.children('td').eq(0).text(data.UserName);
            tr.children('td').eq(1).attr('username-value', data.UserName);
        }
    }
}

function showEditUserDlg(value, event) {
    var userInfo = {
        userName: value.parent().attr('username-value')
    }
    getUserDialog().show(userInfo);
}

function deleteUser(hrefObj, event) {
    value = hrefObj.parent().attr('username-value');
    if (isAdminUserName(value)) {
        notifyErrorMsg('不能删除管理员账号');
        return;
    }
    confirmDialogFunc('确定删除' + value + '吗?', function(e) {
        //向add.ashx请求结果
        $.post('/user/delete', {
                username: value
            },
            //回调函数
            function(data) {
                console.log(data)
                hideConfirmDialog();
                if (!data.Succ) {
                    notifyErrorMsg(data.Msg);
                } else {
                    hrefObj.parent().parent().remove();
                    notifySuccessMsg("删除用户成功");
                }
            },
            //返回类型
            "json"
        );
    });
}