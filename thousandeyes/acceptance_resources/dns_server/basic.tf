data "thousandeyes_agent" "amsterdam" {
  agent_name = "Amsterdam, Netherlands"
}

resource "thousandeyes_alert_rule" "test" {
  rule_name                 = "Custom UAT DNS Server Alert Rule"
  alert_type                = "DNS Server"
  expression                = "((probDetail != \"\") && (Auto(delay >= Medium sensitivity)))"
  minimum_sources           = 1
  rounds_violating_required = 1
  rounds_violating_out_of   = 1
}

resource "thousandeyes_dns_server" "test" {
  test_name      = "User Acceptance Test - DNS Server"
  interval       = 120
  alerts_enabled = true
  domain         = "thousandeyes.com A"

  agents {
    agent_id = data.thousandeyes_agent.amsterdam.agent_id
  }

  alert_rules {
    rule_id = thousandeyes_alert_rule.test.id
  }

  alert_rules {
    rule_id = 921612 #DNS Server Default Alert Rule
  }

  dns_servers {
    server_name = "ns-cloud-d1.googledomains.com"
  }

  dns_servers {
    server_name = "ns-1458.awsdns-54.org"
  }

  dns_servers {
    server_name = "ns-597.awsdns-10.net"
  }
}
