<body style="text-align:center;">
  <table style="max-height:209px;overflow:auto;margin-right:auto;margin-left:auto;" border="1" cellpadding="10">
    <tr>
      <th>vip</th>
      <th>ip:port</th>
      <th>vpnIPs</th>
      <th>routeIPs</th>
      <th>createAt</th>
      <th>lastAccessAt</th>
      <th>transBytes</th>
    </tr>
    {{range .clis}}
    <tr>
        <td align="center">{{ .Vip }}</td>
        <td align="center">{{ .Ip }}</td>
        <td align="center">{{ .VpnIPs }}</td>
        <td align="center">{{ .LanIPs }}</td>
        <td align="center">{{ .CreateTime }}</td>
        <td align="center">{{ .LastAccessTime }}</td>
        <td align="center">{{ .TransBytes }}</td>
    </tr>
    {{end}}
  </table>
</body>
