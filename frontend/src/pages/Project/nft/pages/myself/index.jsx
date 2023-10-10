import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Button, Space, Popconfirm, Avatar, Table, message, Row,Col,Tabs,Card,Collapse } from 'antd'
import { preRouter, STATIC_IMG } from 'src/commons/PRE_ROUTER'
import Header from '../../components/header'
import Footer from '../../components/footer'
import config from 'src/commons/config-hoc';

import './style.less';
@config({
    ajax: true,
    noFrame: true,
    noAuth: true,
  })
export default class nft_detail extends Component {
  state = {
    loading:false,
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
    data:[],
    user:[]
  };
  getWork(id){
    localStorage.setItem('worksId',id)
  }
  getColl(id){
    localStorage.setItem('collectionId',id)
  }
  componentDidMount() {
    this.handleData(1)
  }

  handleData(key){
    this.setState({loading: true})

    const apis = {
      1: 'api/nft/works/findWorksInUser',
      2: 'api/nft/collection/findCollectionsInUser',
      3: 'api/system/user/hisTranInUser',
      4: 'api/nft/works/findFollowsInUser',
      5: 'api/system/user/getInfo'
    }

    this.props.ajax.post(apis[key]).then(res => {
        if (res.status == 'success' && res.statusCode == '200') {
          this.setState({data: res.result})
        } else {
          message.error(res.message);
        }
      }).finally(() => this.setState({ loading: false }));
      
    this.props.ajax.post(apis[5]).then(res => {
      if (res.status == 'success' && res.statusCode == '200') {
        this.setState({user: res.result})
      } else {
        message.error(res.message);
      }
    }).finally(() => this.setState({ loading: false }));
  }
  render() {
    const { data, loading, columns,user } = this.state
    const { TabPane } = Tabs;
    return (
      <div styleName="collection">
        <div styleName='nav'>
          <Header />
        </div>
        <div styleName='banner'>
          <img alt="example" src={STATIC_IMG+"/avatar/collectionbanner.png"} />
          {/* <img alt="example" src={require('../../assets/collectionbanner.png')} /> */}
        </div>
        <div styleName='people'>
          {/* <img alt="example" src={require('../../assets/collectionavatar.png')} /> */}
          <img alt="example" src={STATIC_IMG+user.avatar} />
        </div>
        <div styleName='content'>
          <h2>{user.userName}</h2>
          <p styleName='author'>{user.address}</p>
          <p styleName='date'>{user.createTime}加入</p>
        </div>
        <div styleName='similar'>
        <Tabs defaultActiveKey="1" onChange={key => this.handleData(key)}>
          <TabPane tab={
              <div>
              <img style={{marginTop:'-5px'}} width='15px' heihgt='15px' alt="example" src={require('../../assets/similar.png')} />
              <span>&nbsp;&nbsp;拥有</span></div>
            } key="1">
            <Row gutter={30}>
              {data.map(item=>
                  <Col span={6}>
                      <Card
                          styleName='card'
                          hoverable
                          cover={<Link to={`${preRouter}/nft_detail`} onClick={()=>this.getWork(item.id)}>
                          <img style={{width:'100%',objectFit:'contain'}} alt="example" src={STATIC_IMG + item.dataPath} /></Link>}
                          // <img style={{width:'100%',objectFit:'contain'}} alt="example" src={AJAX_PREFIXCONFIG + item.dataPath} /></Link>}
                          >
                          <Row>
                          <Col span={16}>
                              <p>{item.name}</p>
                              <h3>{item.id}</h3>
                          </Col>
                          <Col span={8} styleName='right'>
                              <p>价格</p>
                              <h3>{item.price}</h3>
                          </Col>
                          </Row>
                          <div style={{textAlign:'center',marginBottom:'15px'}}>
                          {item.state === 0 && item.userId === item.authorId ?
                          <Space>
                          <Link to={{pathname:`${preRouter}/nft_creatNFT`,state:{worksId:item.id}}}>
                              <Button style={{borderRadius:'20px'}} type='primary'>修改</Button>
                          </Link>
                          <Popconfirm style={{position:'absolute'}} title='确定删除该作品吗？' onConfirm={() => { this.handleDelete(item.id) }} okText="确定" cancelText="取消">
                              <Button type='danger' style={{borderRadius:'20px'}}>删除</Button>
                          </Popconfirm>
                          </Space>
                          : item.state === 1 && item.userId === item.authorId ? <Button style={{width:'50%',height:'30px',borderRadius:'20px'}} disabled>已上架</Button>
                          : item.state === 2 && item.userId === item.authorId ? <Button style={{width:'50%',height:'30px',borderRadius:'20px'}} disabled>已售停</Button>
                          : item.state === 3 && item.userId === item.authorId ? <Button style={{width:'50%',height:'30px',borderRadius:'20px'}} disabled>已下架</Button>
                          : item.state === 4 && item.userId === item.authorId ? <Button style={{width:'50%',height:'30px',borderRadius:'20px'}} disabled>已删除</Button>
                          : <Button type='primary' style={{width:'50%',height:'30px',borderRadius:'20px'}} disabled>转移</Button>}
                          </div>
                      </Card>
                  </Col>
              )}
            </Row>
          </TabPane>
          <TabPane tab={
              <div>
              <img style={{marginTop:'-5px'}} width='15px' heihgt='15px' alt="example" src={require('../../assets/creat.png')} />
              <span>&nbsp;&nbsp;创建的合集</span></div>
            } key="2">
            <Row gutter={30}>
            {data?.map(item=>
                <Col span={12}>
                  <Card
                    styleName='card'
                    hoverable
                    cover={<Link to={`${preRouter}/nft_collection`} onClick={()=>this.getColl(item.collectionId)}>
                      <img style={{width:'100%',objectFit:'contain'}} alt="" src={STATIC_IMG + item.coverPath} /></Link>}>
                      {/* <img style={{width:'100%',objectFit:'contain'}} alt="" src={AJAX_PREFIXCONFIG + item.coverPath} /></Link>}> */}
                      <Avatar styleName='img' alt="example" src={STATIC_IMG + item.logoPath} />
                      {/* <Avatar styleName='img' alt="example" src={AJAX_PREFIXCONFIG + item.logoPath} /> */}
                      <h3>{item.name}</h3>
                      <p>作者<span style={{color:'#2345A7'}}>{item.authorName}</span></p>
                      {item.details ? <div className="braft-output-content" dangerouslySetInnerHTML = {{ __html:decodeURIComponent(item.details) }}></div>
                      : <div>暂无简介</div>}

                    {item.state === 0 && item.userId === item.authorId ?
                    <Link to={{pathname:`${preRouter}/nft_creatCollection`,state:{collectionId:item.collectionId}}}>
                      <Button style={{width:'50%',height:'30px',borderRadius:'20px'}} type='primary'>修改</Button>
                    </Link>
                    : item.state === 1 && item.userId === item.authorId ? <Button disabled>已上架</Button>
                    : item.state === 2 && item.userId === item.authorId ? <Button disabled>已售停</Button>
                    : item.state === 3 && item.userId === item.authorId ? <Button disabled>已下架</Button>
                    : item.state === 4 && item.userId === item.authorId ? <Button disabled>已删除</Button>
                    : <Button disabled>详情</Button>}
                  </Card>
                </Col>
              )}
            </Row>
          </TabPane>
          <TabPane tab={
            <div>
              <img style={{marginTop:'-5px'}} width='15px' heihgt='15px' alt="example" src={require('../../assets/out.png')} />
              <span>&nbsp;&nbsp;交易历史</span></div>
            } key="3">
            <Table
              loading={loading}
              columns={columns}
              dataSource={data}
              pagination={false}
              rowKey="key"
            />
          </TabPane>
          <TabPane tab={
            <div>
              <img style={{marginTop:'-5px'}} width='15px' heihgt='15px' alt="example" src={require('../../assets/star.png')} />
              <span>&nbsp;&nbsp;关注</span>
            </div>} key="4">
            <Row gutter={30}>
              {data.map(item=>
                  <Col span={6}>
                      <Card
                          styleName='card'
                          hoverable
                          cover={<Link to={`${preRouter}/nft_detail`} onClick={()=>this.getWork(item.id)}>
                          <img style={{width:'100%',objectFit:'contain'}} alt="example" src={STATIC_IMG + item.dataPath} /></Link>}
                          // <img style={{width:'100%',objectFit:'contain'}} alt="example" src={AJAX_PREFIXCONFIG + item.dataPath} /></Link>}
                          >
                          <Row>
                          <Col span={16}>
                              <p>{item.name}</p>
                              <h3>{item.id}</h3>
                          </Col>
                          <Col span={8} styleName='right'>
                              <p>价格</p>
                              <h3>{item.price}</h3>
                          </Col>
                          </Row>
                          <div style={{textAlign:'center',marginBottom:'15px'}}>
                          {item.state === 0 && item.userId === item.authorId ?
                          <Space>
                          <Link to={{pathname:`${preRouter}/nft_creatNFT`,state:{worksId:item.id}}}>
                              <Button style={{borderRadius:'20px'}} type='primary'>修改</Button>
                          </Link>
                          <Popconfirm style={{position:'absolute'}} title='确定删除该作品吗？' onConfirm={() => { this.handleDelete(item.id) }} okText="确定" cancelText="取消">
                              <Button type='danger' style={{borderRadius:'20px'}}>删除</Button>
                          </Popconfirm>
                          </Space>
                          : item.state === 1 && item.userId === item.authorId ? <Button style={{width:'50%',height:'30px',borderRadius:'20px'}} disabled>已上架</Button>
                          : item.state === 2 && item.userId === item.authorId ? <Button style={{width:'50%',height:'30px',borderRadius:'20px'}} disabled>已售停</Button>
                          : item.state === 3 && item.userId === item.authorId ? <Button style={{width:'50%',height:'30px',borderRadius:'20px'}} disabled>已下架</Button>
                          : item.state === 4 && item.userId === item.authorId ? <Button style={{width:'50%',height:'30px',borderRadius:'20px'}} disabled>已删除</Button>
                          : <Button type='primary' style={{width:'50%',height:'30px',borderRadius:'20px'}} disabled>详情</Button>}
                          </div>
                      </Card>
                  </Col>
              )}
            </Row>
          </TabPane>
        </Tabs>
        </div>
        <div styleName='footer'>
          <Footer />
        </div>
      </div>
    )
  }
}
