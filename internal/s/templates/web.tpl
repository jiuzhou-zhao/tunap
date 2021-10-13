<body style="text-align:center;">
  <table style="max-height:209px;overflow:auto;margin-right:auto;margin-left:auto;" border="1">
    <tr>
      <th>vip</th>
      <th>ip:port</th>
      <th>vpnIPs</th>
      <th>routeIPs</th>
    </tr>
    {{range .clis}}
    <tr>
        <td>{{ .Vip }}</td>
        <td>{{ .Ip }}</td>
        <td>{{ .VpnIPs }}</td>
        <td>{{ .LanIPs }}</td>
    </tr>
    {{end}}
  </table>
</body>
