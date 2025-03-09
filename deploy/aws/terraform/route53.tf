data "aws_route53_zone" "rootdns" {
  name = var.dns_root
}

resource "aws_route53_record" "env_record" {
  name    = "${var.dns_subdomain}.${var.dns_root}"
  type    = "A"
  zone_id = data.aws_route53_zone.rootdns.zone_id
  alias {
    evaluate_target_health = true
    name                   = aws_cloudfront_distribution.website.domain_name
    zone_id                = aws_cloudfront_distribution.website.hosted_zone_id
  }
}