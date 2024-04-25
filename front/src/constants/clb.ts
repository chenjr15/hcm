// 负载均衡-路由组件名称
export enum LBRouteName {
  allLbs = 'all-lbs-manager',
  lb = 'specific-lb-manager',
  listener = 'specific-listener-manager',
  domain = 'specific-domain-manager',
  allTgs = 'all-tgs-manager',
  tg = 'specific-tg-manager',
}
// 负载均衡-路由组件名称映射
export const LB_ROUTE_NAME_MAP = {
  all: 'all-lbs-manager',
  lb: 'specific-lb-manager',
  listener: 'specific-listener-manager',
  domain: 'specific-domain-manager',
};

// 网络类型
export const LOAD_BALANCER_TYPE = [
  {
    label: '公网',
    value: 'OPEN',
  },
  {
    label: '内网',
    value: 'INTERNAL',
  },
];
// IP版本
export const ADDRESS_IP_VERSION = [
  {
    label: 'IPV4',
    value: 'IPV4',
  },
  {
    label: 'IPV6',
    value: 'IPv6FullChain',
  },
  {
    label: 'IPV6 NAT64',
    value: 'IPV6',
    isDisabled: (region: string) => !WHITE_LIST_REGION_IPV6_NAT64.includes(region),
  },
];
// 可用区类型
export const ZONE_TYPE = [
  {
    label: '单可用区',
    value: 'single',
  },
  {
    label: '主备可用区',
    value: 'primaryStand',
    isDisabled: (region: string) => !WHITE_LIST_REGION_PRIMARY_STAND_ZONE.includes(region),
  },
];
// 网络计费模式
export const INTERNET_CHARGE_TYPE = [
  {
    label: '包月',
    value: undefined,
  },
  {
    label: '按流量',
    value: 'TRAFFIC_POSTPAID_BY_HOUR',
  },
  {
    label: '按带宽',
    value: 'BANDWIDTH_POSTPAID_BY_HOUR',
  },
  // {
  //   label: '共享带宽包',
  //   value: 'BANDWIDTH_PACKAGE',
  // },
];

// 支持IPv6 NAT64的地域
export const WHITE_LIST_REGION_IPV6_NAT64 = ['ap-beijing', 'ap-shanghai', 'ap-guangzhou'];
// 支持主备可用区的地域
export const WHITE_LIST_REGION_PRIMARY_STAND_ZONE = [
  'ap-guangzhou',
  'ap-shanghai',
  'ap-nanjing',
  'ap-beijing',
  'ap-hongkong',
  'ap-seoul',
];

// 会话类型映射
export const SESSION_TYPE_MAP = {
  NORMAL: '基于源 IP ',
  QUIC_CID: '基于源端口',
};

// 证书认证方式映射
export const SSL_MODE_MAP = {
  UNIDIRECTIONAL: '单向认证',
  MUTUAL: '双向认证',
};

// 均衡方式映射
export const SCHEDULER_MAP = {
  WRR: '按权重轮询',
  LEAST_CONN: '最小连接数',
  IP_HASH: 'IP Hash',
};

// 传输层协议, 如 TCP, UDP
export const TRANSPORT_LAYER_LIST = ['TCP', 'UDP'];
// 应用层协议, 如 HTTP, HTTPS
export const APPLICATION_LAYER_LIST = ['HTTP', 'HTTPS'];

// 负载均衡网络类型映射
export const LB_NETWORK_TYPE_MAP = {
  OPEN: '公网',
  INTERNAL: '内网',
};
