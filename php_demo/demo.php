<?php

// 接口请求域名
$main_url = 'https://api.salesmartly.com';

// 项目API Token
$api_token = 'api_token';

// 项目Id
$project_id = 'project_id';

// 生成MD5签名
function generate_md5_signature($api_token, $params) {
    ksort($params); // 对参数进行排序
    $sign_string = $api_token;
    foreach ($params as $key => $value) {
        $sign_string .= "&" . $key . "=" . $value;
    }
    echo "签名字符串: " . $sign_string . "\n";
    return md5($sign_string);
}

// 发送 GET 请求
function send_get_request($api_url, $params = [], $headers = []) {
    global $main_url;

    $url = $main_url . $api_url;

    // 将参数附加到URL
    if (!empty($params)) {
        $url .= '?' . http_build_query($params);
    }

    // 初始化cURL
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);

    // 执行请求
    $response = curl_exec($ch);

    // 检查请求是否成功
    if (curl_errno($ch)) {
        echo "GET 请求失败: " . curl_error($ch);
    } else {
        $http_code = curl_getinfo($ch, CURLINFO_HTTP_CODE);
        if ($http_code == 200) {
            $data = json_decode($response, true);
            echo "GET 请求成功，响应数据：\n";
            print_r($data);
        } else {
            echo "GET 请求失败，HTTP状态码: " . $http_code;
        }
    }

    // 关闭cURL
    curl_close($ch);
}

// 发送 POST 请求
function send_post_request($api_url, $payload = [], $headers = []) {
    global $main_url;

    $url = $main_url . $api_url;

    // 初始化cURL
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_POST, true);
    curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($payload));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);

    // 执行请求
    $response = curl_exec($ch);

    // 检查请求是否成功
    if (curl_errno($ch)) {
        echo "POST 请求失败: " . curl_error($ch);
    } else {
        $http_code = curl_getinfo($ch, CURLINFO_HTTP_CODE);
        if ($http_code == 200) {
            $data = json_decode($response, true);
            echo "POST 请求成功，响应数据：\n";
            print_r($data);
        } else {
            echo "POST 请求失败，HTTP状态码: " . $http_code;
        }
    }

    // 关闭cURL
    curl_close($ch);
}

// 主程序
//if (__FILE__ == $_SERVER['SCRIPT_FILENAME']) {
    // 定义请求的接口及参数
    $get_url = '/api/chat-user/get-contact-list';
    $get_params = [
        'project_id' => $project_id,
        'updated_time' => '{"start":1680000000,"end":1814027206}'
    ];

    // 参数加密出签名
    $get_sign = generate_md5_signature($api_token, $get_params);
    $get_header = [
        'external-sign: ' . $get_sign
    ];

    // 发送GET请求
    send_get_request($get_url, $get_params, $get_header);
//}