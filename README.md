## tunap 内网穿透


## 修复windows ICMP无回显

* 以管理员身份运行下面两条命令：
```bash
netsh advfirewall firewall add rule name="ICMP Allow incoming V4 echo request“ protocol=icmpv4:8,any dir=in action=allow

netsh advfirewall firewall add rule name="ICMP Allow incoming V6 echo request” protocol=icmpv6:8,any dir=in action=allow
```

* 或者添加靜態：

```bash
route -p add 目的地址 mask 子网掩码 网关地址
```
它的意思是，要想找到“目的地址”，就要通过“网关地址”里面找。-p是永久有效的意思。

添加完毕后可用 route print 查看是否添加成功

* 另外两种方法：

1、关掉自带防火墙；（不建议，毕竟防火墙最后一道屏障）

2、ICMPv4 两个都点启用。点击“入站规则”，在右侧“入站规则”找到“文件和打印机共享（回显请求）”注意是ICMPv4。然后启用规则

