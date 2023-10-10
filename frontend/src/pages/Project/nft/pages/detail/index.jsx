import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Row,Col,Table, Button, message,Card,Collapse } from 'antd'
import { preRouter,AJAX_PREFIXCONFIG,STATIC_IMG } from 'src/commons/PRE_ROUTER'
import { DownOutlined } from '@ant-design/icons';
import Header from '../../components/header'
import Footer from '../../components/footer'
import LineEcharts from './components/lineEcharts';
import config from 'src/commons/config-hoc';
import Shelves from './components/Shelves'
import './style.less';

const localStorage = window.localStorage;
const CODEMAP = JSON.parse(localStorage.getItem('login-user')).codeMap;

@config({
  path: '/nft_detail',
  ajax: true,
  noFrame: true,
  noAuth: true,
})
export default class nft_detail extends Component {
  state = {
    loading:false,
    avaData:[],
    userData:[],
    visible:false,
    columns:[
      {
        title: '事件', dataIndex: 'tradeType', key: 'tradeType',
        render: (value, record) => {
            const { tradeType } = record;
            return (
              <div>
                {tradeType === 6 ? <span><img alt="example" src={require('../../assets/out.png')} />&nbsp;&nbsp;转移</span>
                : <span><img alt="example" src={require('../../assets/sell.png')} />&nbsp;&nbsp;交易</span>}
              </div>
            )
        },
      },
      {
        title: '价格', dataIndex: 'price', key: 'price',
      },
      {
        title: '卖方', dataIndex: 'fromUser', key: 'fromUser',
      },
      {
        title: '买方', dataIndex: 'toUser', key: 'toUser',
      },
      {
        title: '时间', dataIndex: 'createTime', key: 'createTime',
      },
    ],
    data:[
      /* {text:'转移',a:'1',b:'HYAGJ',c:'CPPPP',date:'1天前'},
      {text:'售卖',a:'1',b:'HYAGJ',c:'CPPPP',date:'1天前'},
      {text:'转移',a:'1',b:'HYAGJ',c:'CPPPP',date:'1天前'},
      {text:'转移',a:'1',b:'HYAGJ',c:'CPPPP',date:'1天前'},
      {text:'售卖',a:'1',b:'HYAGJ',c:'CPPPP',date:'1天前'}, */
    ],
    collectStatus: 0
  };

  componentDidMount() {
    this.handleAvatar()
    this.handleUser()
    this.fetchCollect()
    this.fetchTransaction()
  }

  fetchTransaction() {
    const worksId = localStorage.getItem('worksId')
    this.props.ajax.post(`api/nft/works/hisTrans?worksId=${worksId}`).then(res => {
        if (res.status == 'success' && res.statusCode == '200') {
            this.setState({data:res.result})
        } else {
        message.error(res.message);
        }
    }).finally(() => this.setState({ loading: false }));
  }

  fetchCollect() {
    const worksId = localStorage.getItem('worksId')
    this.props.ajax.post(`api/system/user/followStatus?worksId=${worksId}`).then(res => {
        if (res.status == 'success' && res.statusCode == '200') {
            this.setState({collectStatus:res.result})
        } else {
        message.error(res.message);
        }
    }).finally(() => this.setState({ loading: false }));
  }

  //获取头像数据
  handleAvatar(){
    const worksId = localStorage.getItem('worksId')
    this.setState({loading:true})
    this.props.ajax.post(`${CODEMAP.nft_mobile_works_buyWorksInfo}?worksId=${worksId}`).then(res => {
        if (res.status == 'success' && res.statusCode == '200') {
        this.setState({avaData:res.result})
        } else {
        message.error(res.message);
        }
    }).finally(() => this.setState({ loading: false }));
  }

  //获取持有数据
  handleUser(){
    const worksId = localStorage.getItem('worksId')
    this.setState({loading:true})
    this.props.ajax.post(`${CODEMAP.nft_mobile_works_buyWorks}?worksId=${worksId}`).then(res => {
        if (res.status == 'success' && res.statusCode == '200') {
        this.setState({userData:res.result})
        } else {
        message.error(res.message);
        }
    }).finally(() => this.setState({ loading: false }));
  }

  collect() {
    const worksId = localStorage.getItem('worksId')
    const { collectStatus } = this.state

    this.setState({collectStatus: collectStatus === 1 ? 0 : 1 })

    setTimeout(() => {
        this.props.ajax.post(`api/system/user/fllowWorks?worksId=${worksId}`).then(res => {
            if (res.statusCode !== 200) {
                message.error(res.message);
            }
        }).finally(() => this.setState({ loading: false }));
    })
  }

  buyNftWork() {
    const worksId = localStorage.getItem('worksId')
    this.props.ajax.post(`api/nft/works/buyNftWork?worksId=${worksId}`).then(res => {
        if (res.status == 'success' && res.statusCode == '200') {
          window.location.reload()
        } else {
        message.error(res.message);
        }
    }).finally(() => this.setState({ loading: false }));
  }


  render() {
    const { Panel } = Collapse;
    const { loading,columns,data,avaData,userData,visible,collectStatus } = this.state
    return (
      <div styleName="detail">
        <div styleName='nav'>
          <Header />
        </div>
        <div styleName='content'>
          <Row gutter={20}>
            <Col span={10}>
              <img src={STATIC_IMG + userData?.logoPath} alt="" />
              {/* <img src={AJAX_PREFIXCONFIG + userData?.logoPath} alt="" /> */}
            <br /><br />
            <Card styleName='author' title={<div>
              <img alt="example" src={require('../../assets/author.png')} />
              <span>&nbsp;&nbsp;作者</span></div>}>
              <div styleName='name'>{avaData?.author}<span>创作</span></div>
              <Collapse bordered={false} accordion>
                <Panel showArrow={false} header={<p styleName='edu'><img alt="example" src={require('../../assets/detail1.png')} />&nbsp;&nbsp;链上信息<DownOutlined style={{float:'right',marginTop:'5px'}}/></p>} key="1">
                  <div className="braft-output-content" dangerouslySetInnerHTML = {{ __html:decodeURIComponent(avaData?.detail) }}></div>
                </Panel>
                <Panel showArrow={false} header={<p styleName='edu'><img alt="example" src={require('../../assets/detail2.png')} />&nbsp;&nbsp;合集简介<DownOutlined style={{float:'right',marginTop:'5px'}}/></p>} key="2">
                <div className="braft-output-content" dangerouslySetInnerHTML = {{ __html:decodeURIComponent(avaData?.collectionDetail) }}></div>
                </Panel>
                <Panel showArrow={false} header={<p styleName='edu'><img alt="example" src={require('../../assets/detail3.png')} />&nbsp;&nbsp;链上信息<DownOutlined style={{float:'right',marginTop:'5px'}}/></p>} key="3">
                <div className="braft-output-content" style={{overflow:'hidden',textOverflow:'ellipsis'}} dangerouslySetInnerHTML = {{ __html:decodeURIComponent(avaData?.chainInfo) }}></div>
                </Panel>
              </Collapse>
            </Card>
            </Col>
            <Col span={14}>
              <Row>
                <Col span={20}>
                  <div styleName='bored'>
                  <h3>{userData?.worksName}</h3>
                  <h2>#{userData?.worksNum}</h2>
                    <p>购买者<span>{userData?.buyer}</span></p>
                  </div>
                </Col>
                <Col span={4}>
                  <div styleName='star' onClick={() => this.collect()}>
                    {/* <img alt="example" src={require('../../assets/star.png')} /> */}
                    <span style={{ color: collectStatus ? '#ffb300' : '#ccc' }}>★</span>
                  </div>
                </Col>
              </Row>
              <Card>
                <div styleName='price'>
                  <h3>剩余数量：{userData?.remainingNum}</h3>
                  <p>目前价格</p>
                  <h2>&nbsp;￥{userData?.price}</h2>
                  {userData?.state == 0 && userData?.authorId == userData?.userId ? <Button type='primary' onClick={()=>this.setState({visible:true})}>上架</Button>
                  : userData?.state == 0 && userData?.authorId != userData?.userId ? <Button type='primary' disabled>未上架</Button>
                  : userData?.state == 1 && userData?.authorId == userData?.userId ? <Button type='primary' disabled>下架</Button>
                  : userData?.state == 1 && userData?.authorId != userData?.userId ? <Button type='primary' onClick={()=>this.buyNftWork()}>购买</Button>
                  : userData?.state == 2 && userData?.authorId == userData?.userId ? <Button type='primary' disabled>售停</Button>
                  : userData?.state == 2 && userData?.authorId != userData?.userId ? <Button type='primary' disabled>售停</Button>
                  : userData?.state == 3 && userData?.authorId == userData?.userId ? <Button type='primary' disabled>已下架</Button>
                  : userData?.state == 3 && userData?.authorId != userData?.userId ? <Button type='primary' disabled>已下架</Button>
                  : userData?.state == 4 && userData?.authorId == userData?.userId ? <Button type='primary' disabled>已删除</Button>
                  : userData?.state == 4 && userData?.authorId != userData?.userId ? <Button type='primary' disabled>已删除</Button>
                  : <Button type='primary'></Button>}
                </div>
              </Card>
              <br />
              <Card styleName='history' title={<div>
              <img alt="example" src={require('../../assets/author.png')} />
              <span>&nbsp;&nbsp;历史价格</span></div>}>
                <LineEcharts />
              </Card>
            </Col>
          </Row>
        </div>
        <div styleName='trading'>
          <Card title={<div>
            <img alt="example" src={require('../../assets/out.png')} />
            <span>&nbsp;&nbsp;交易信息</span></div>}>
            <Table
              loading={loading}
              columns={columns}
              dataSource={data}
              pagination={false}
              rowKey="key"
            />
          </Card>
        </div>
        <div styleName='similar'>
          <Card title={<div>
            <img width='20px' heihgt='20px' alt="example" src={require('../../assets/similar.png')} />
            <span>&nbsp;&nbsp;更多相似内容</span></div>}>
            <Row gutter={30}>
              <Col span={6}>
                <Card
                    hoverable
                    cover={<img alt="example" src={require('../../assets/detailbanner1.png')} />}
                  >
                    <Row>
                      <Col span={16}>
                        <p>Bored Ape Yac..</p>
                        <h3>4578</h3>
                      </Col>
                      <Col span={8} styleName='right'>
                        <p>价格</p>
                        <h3 styleName='price'><span></span>12.36</h3>
                      </Col>
                    </Row>
                  </Card>
              </Col>
              <Col span={6}>
                <Card
                    hoverable
                    cover={<img alt="example" src={require('../../assets/detailbanner2.png')} />}
                  >
                    <Row>
                      <Col span={16}>
                        <p>Bored Ape Yac..</p>
                        <h3>4578</h3>
                      </Col>
                      <Col span={8} styleName='right'>
                        <p>价格</p>
                        <h3 styleName='price'><span></span>12.36</h3>
                      </Col>
                    </Row>
                  </Card>
              </Col>
              <Col span={6}>
                <Card
                    hoverable
                    cover={<img alt="example" src={require('../../assets/detailbanner3.png')} />}
                  >
                    <Row>
                      <Col span={16}>
                        <p>Bored Ape Yac..</p>
                        <h3>4578</h3>
                      </Col>
                      <Col span={8} styleName='right'>
                        <p>价格</p>
                        <h3 styleName='price'><span></span>12.36</h3>
                      </Col>
                    </Row>
                  </Card>
              </Col>
              <Col span={6}>
                <Card
                    hoverable
                    cover={<img alt="example" src={require('../../assets/detailbanner4.png')} />}
                  >
                    <Row>
                      <Col span={16}>
                        <p>Bored Ape Yac..</p>
                        <h3>4578</h3>
                      </Col>
                      <Col span={8} styleName='right'>
                        <p>价格</p>
                        <h3 styleName='price'><span></span>12.36</h3>
                      </Col>
                    </Row>
                  </Card>
              </Col>
            </Row>
          </Card>
        </div>
        <Shelves
          visible={visible}
          onOk={() => {this.setState({visible:false});this.handleUser()}}
          onCancel={() => this.setState({visible:false})}
        />
        <div styleName='footer'>
          <Footer />
        </div>
      </div>
    )
  }
}
