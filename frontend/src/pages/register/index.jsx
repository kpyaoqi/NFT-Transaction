import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Input, Radio, Button, Upload, Avatar, Form, Row, Col, message } from "antd";
import { LoadingOutlined, UploadOutlined, LockOutlined, UserOutlined, MobileOutlined, MailOutlined } from "@ant-design/icons";
import { setLoginUser, toHome, locationHref,toLogin } from "src/commons";
import { preRouter } from "src/commons/PRE_ROUTER";
import config from "src/commons/config-hoc";
import Banner from "./banner/index";
import "./style.less";
import { AJAX_PREFIXCONFIG } from "src/commons/PRE_ROUTER";
function beforeUpload(file) {
    const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
    if (!isJpgOrPng) {
        message.error('You can only upload JPG/PNG file!');
    }
    const isLt2M = file.size / 1024 / 1024 < 2;
    if (!isLt2M) {
        message.error('Image must smaller than 2MB!');
    }
    return isJpgOrPng && isLt2M;
}

function getBase64(img, callback) {
    const reader = new FileReader();
    reader.addEventListener('load', () => callback(reader.result));
    reader.readAsDataURL(img);
}

@config({
    path: "/login",
    ajax: true,
    noFrame: true,
    noAuth: true,
})
export default class Login extends Component {
    state = {
        imageUrl: null,
        loading: false,
        message: "",
        avatar: '',
        isMount: false,
        captchaData: {},
        showCaptcha: false,
    };

    componentDidMount() {
        // 开发时方便测试，填写表单
        if (process.env.NODE_ENV === "development" || process.env.PREVIEW) {
            this.form.setFieldsValue({ username: "", password: "" });
        }
        setTimeout(() => this.setState({ isMount: true }), 300);
    }

    //获取注册验证码
    getCaptcha = () => {
        this.props.ajax
            .post(`api/getCaptcha`, null, { errorTip: true, noEmpty: true })
            .then((res) => {
                this.setState({
                    captchaData: res.result,
                });
            });
    };

    //注册提交操作
    handleSubmit = (values) => {
        if (this.state.loading) return;

        const {
            userName,
            account,
            password,
            sex,
            mobile,
            email,
        } = values;
        const { avatar } = this.state;
        const params = {
            userName,
            account,
            password,
            sex,
            mobile,
            avatar,
            email
        };
        this.setState({ loading: true, message: "" });
        this.props.ajax
            .post(`api/system/user/save`, params)
            .then((res) => {
                if (res.message === "操作成功") {
                    toLogin();
                }
                this.setState({ loading: false, message: res.message });

                /* if (res.status == "success" && res.statusCode == "200") {
                    toHome();
                } else {
                    if (res.result.showCaptcha) {
                        this.setState({ showCaptcha: true }, () => {
                            this.getCaptcha();
                        });
                    }

                    this.setState({ loading: false, message: res.message });
                } */
            })
            .catch((error) => {
                this.setState({ message: error });
            })
            .finally(() => this.setState({ loading: false }));
    };

    handleChange = info => {
        if (info.file.status === 'uploading') {
            this.setState({ loading: true });
            return;
        }
        if (info.file.status === 'done') {

            this.setState({avatar:info.file.response.result.fileURL})
            this.form.setFieldsValue({ avatar: info.file.response.result.fileURL});

            // Get this url from response in real world.
            getBase64(info.file.originFileObj, imageUrl =>
            this.setState({
                imageUrl:imageUrl,
                loading: false,
            })

            );
        }
    };

    render() {
        const {
            imageUrl,
            loading,
            message,
            isMount,
            captchaData,
            showCaptcha
        } = this.state;

        const uploadButton = (
            <div style={{textAlign:'center',lineHeight:'150px'}}>
              <div style={{ marginTop: 8 }}>{loading ? <LoadingOutlined /> : <UploadOutlined style={{fontSize:'40px',opacity:'0.5'}} />}</div>
            </div>
          );

        const formItemStyleName = isMount ? "form-item active" : "form-item";

        return (
            <div styleName="root">
                <Helmet title="欢迎登陆" />
                <div styleName="banner">
                    <Banner />
                </div>
                <div styleName="box">
                    <Form
                        ref={(form) => (this.form = form)}
                        name="login"
                        className="inputLine"
                        onFinish={this.handleSubmit}
                    >
                        <div styleName={formItemStyleName}>
                            <div styleName="header">欢迎注册</div>
                        </div>
                        <div styleName={formItemStyleName}>
                            <Form.Item
                                name="userName"
                                rules={[
                                    { required: true, message: "请输入用户名" },
                                ]}
                            >
                                <Input
                                    allowClear
                                    autoFocus
                                    prefix={
                                        <UserOutlined className="site-form-item-icon" />
                                    }
                                    placeholder="用户名"
                                />
                            </Form.Item>
                        </div>
                        <div styleName={formItemStyleName}>
                            <Form.Item
                                name="account"
                                rules={[
                                    { required: true, message: "请输入账号" },
                                ]}
                            >
                                <Input
                                    allowClear
                                    autoFocus
                                    prefix={
                                        <UserOutlined className="site-form-item-icon" />
                                    }
                                    placeholder="账号"
                                />
                            </Form.Item>
                        </div>
                        <div styleName={formItemStyleName}>
                            <Form.Item
                                name="password"
                                rules={[
                                    { required: true, message: "请输入密码" },
                                ]}
                            >
                                <Input.Password
                                    prefix={
                                        <LockOutlined className="site-form-item-icon" />
                                    }
                                    placeholder="密码"
                                />
                            </Form.Item>
                        </div>
                        <div styleName={formItemStyleName}>
                            <Form.Item
                                label="性别"
                                name="sex"
                                rules={[
                                    { required: true, message: "请选择性别" },
                                ]}
                            >
                                <Radio.Group>
                                    <Radio value="男">男</Radio>
                                    <Radio value="女">女</Radio>
                                </Radio.Group>
                            </Form.Item>
                        </div>
                        <div styleName={formItemStyleName}>
                            <Form.Item
                                name="mobile"
                                rules={[
                                    { required: true, message: "请输入手机号码" },
                                ]}
                            >
                                <Input
                                    allowClear
                                    autoFocus
                                    prefix={
                                        <MobileOutlined className="site-form-item-icon" />
                                    }
                                    placeholder="手机号码"
                                />
                            </Form.Item>
                        </div>
                        <div styleName={formItemStyleName}>
                            <Form.Item
                                styleName="form-item-upload"
                                label="头像照片"
                                name="avatar"
                                rules={[
                                    { required: true, message: "请上传头像" },
                                ]}
                                >
                                <p>建议使用150*150等比例图片</p>
                                <Upload
                                    name="file"
                                    showUploadList={false}
                                    action={`${AJAX_PREFIXCONFIG}/attachment/upload`}
                                    beforeUpload={beforeUpload}
                                    onChange={this.handleChange}
                                    headers={{contentType: 'multipart/form-data'}}
                                    // headers={{contentType: 'multipart/form-data',jwttoken:`${JSON.parse(sessionStorage.getItem('login-user')).token}`}}
                                    data={{type:0}}
                                >
                                    {imageUrl
                                        ? <div style={{width:'150px',height:'150px',borderRadius:'50%',overflow:'hidden'}}><Avatar src={imageUrl} alt="avatar" style={{ width: '100%',height:'100%'}} /></div>
                                        : <div style={{width:'150px',height:'150px',border:'2px dashed #6B6E71',borderRadius:'50%'}}>{uploadButton}</div>}
                                </Upload>
                            </Form.Item>
                        </div>
                        <div styleName={formItemStyleName}>
                            <Form.Item
                                name="email"
                            >
                                <Input
                                    allowClear
                                    autoFocus
                                    prefix={
                                        <MailOutlined className="site-form-item-icon" />
                                    }
                                    placeholder="邮箱"
                                />
                            </Form.Item>
                        </div>
                        {showCaptcha ? (
                            <div styleName={formItemStyleName}>
                                <Row gutter={24}>
                                    <Col span={16}>
                                        <Form.Item
                                            name="captcha"
                                            rules={[
                                                {
                                                    required: true,
                                                    message: "请输入验证码",
                                                },
                                            ]}
                                        >
                                            <Input
                                                prefix={
                                                    <LockOutlined className="site-form-item-icon" />
                                                }
                                                placeholder="请输入验证码"
                                            />
                                        </Form.Item>
                                    </Col>
                                    <Col span={8}>
                                        <img
                                            onClick={() => {
                                                this.getCaptcha();
                                            }}
                                            styleName="captchaStyle"
                                            src={captchaData.imageBase64}
                                            alt=""
                                        />
                                    </Col>
                                </Row>
                            </div>
                        ) : null}
                        <div styleName={formItemStyleName}>
                            <Form.Item
                                shouldUpdate={true}
                                style={{ marginBottom: 0 }}
                            >
                                {() => (
                                    <Button
                                        styleName="submit-btn"
                                        loading={loading}
                                        type="primary"
                                        htmlType="submit"
                                        disabled={
                                            !this.form?.isFieldsTouched(true) ||
                                            this.form
                                                ?.getFieldsError()
                                                .filter(
                                                    ({ errors }) =>
                                                        errors.length
                                                ).length
                                        }
                                    >
                                        注册
                                    </Button>
                                )}
                            </Form.Item>
                        </div>
                    </Form>
                    <div styleName="error-tip">{message}</div>
                </div>
            </div>
        );
    }
}
