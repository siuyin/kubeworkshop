# Instumenting an application with Open Telemetry Custom Spans

## Abstract
When OpenTelemetry eBPF Instrumentation (OBI) (https://opentelemetry.io/docs/zero-code/obi/) is installed in a kubernetes cluster,
it can automatically instrument an application.
However it cannot know about your custom business logic.

Auto instumentation will trace a call to /api/v1/payment 
and will create a span when a http call is made to a payment processor, say, a credit card processor,
but it cannot create spans for, say, database lookup of the user's loyalty points or credit score.
