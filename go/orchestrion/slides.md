---
marp: true
paginate: true
---

# Orchestrion: lightning demo

### Automatic compile-time instrumentation of Go code

Dario Castañé (he/him; github: darccio, bsky: dario.cat, fediverse&email: d@rio.hn)
GoLab 2024

<!--
DataDog: dd-trace-go maintainer

In the next minutes, we'll go from a basic chi server without any dd-trace-go code to the exactly chi server
without any dd-trace-go code BUT sending traces and profiles
to DataDog platfor.
-->

---

<style>
img[alt~="center"] {
  display: block;
  margin: 0 auto;
}
</style>

![center](./golab2024.jpeg)

---

# The code

---

# Orchestrion

[github.com/DataDog/orchestrion](github.com/DataDog/orchestrion)

<!--
cat go.mod
orchestrion pin
cat go.mod
-->

---

# Demo time!

- [Tracing](https://app.datadoghq.eu/apm/traces?query=%40_trace_root%3A1%20service%3Agolab2024&agg_m=count&agg_m_source=base&agg_t=count&cols=core_service%2Ccore_resource_name%2Clog_duration%2Clog_http.method%2Clog_http.status_code&fromUser=false&historicalData=false&messageDisplay=inline&query_translation_version=v0&serviceName=golab2024&sort=desc&spanType=trace-root&storage=hot&view=spans&paused=false)
- [Profiling](https://app.datadoghq.eu/profiling/explorer?query=service%3Agolab2024%20host%3ACOMP-WDWT6G66NH&agg_m=count&agg_m_source=base&agg_t=count&fromUser=true&my_code=disabled&refresh_mode=paused&viz=flame_graph&live=true)

---

# Do you want to know more?

Julio Guerra - GopherCon Europe 2021: [youtu.be/Uk1hscXhlY0](https://youtu.be/Uk1hscXhlY0)

Jon Bodner - GopherCon 2023: [youtu.be/5l-W7vPSbuc](https://youtu.be/5l-W7vPSbuc)

Me (again): bsky: dario.cat, fediverse&email: d@rio.hn

These slides: [github.com/darccio/talks/tree/main/go/orchestrion](https://github.com/darccio/talks/tree/main/go/orchestrion)

![w:300 h:300 center](./qr.png)