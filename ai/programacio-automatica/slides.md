---
marp: true
theme: gaia
paginate: true
---

# De filadors de codi a teixidors de sistemes

La revolució industrial del programari

<!--
[EXORDIO — 1 minut]

Començar directe, sense salutació llarga.

"Fa uns mesos, un company em va demanar ajuda amb un mòdul en Rust. Jo sóc gopher de tota la vida — porto anys mantenint llibreries Go. Rust no és el meu llenguatge. Però vaig obrir Claude Code, li vaig donar context sobre el projecte, les restriccions, el que havia de fer... i en dues hores teníem un PR obert. Funcional. Amb tests.

Aquell mes vaig fer el doble de feina. Literalment: dos workstreams en paral·lel, un en Go i un en Rust. El cost? Uns 500 dòlars en tokens.

No us explico això per fer-me el llest. Us ho explico perquè aquell moment em va obligar a repensar què significa ser enginyer de programari avui."
-->

---

## Dario Castañé

Senior Software Engineer @ Datadog

Mantenidor de `dd-trace-go` i `Orchestrion`

Mantenidor de `mergo` — +30.000 repositoris

<!--
[PRESENTACIÓ — 30 segons]

"Em dic Dario Castañé. Treballo a Datadog com a Senior Software Engineer. Mantinc dd-trace-go i Orchestrion, que és la instrumentació automàtica de Go, i també mantinc mergo, una llibreria Go que fan servir més de 30.000 repositoris. Sóc gopher, i això és rellevant per al que us explicaré."
-->

---

# Abans de començar

Dues coses que **no** faré avui:

1. No us diré si la IA és bona o dolenta
2. No us donaré una llista d'eines ni de tècniques

El que sí faré:

Proposar un **marc mental** per entendre el moment

<!--
[DISCLAIMERS — 2 minuts]

"Abans de res, dues coses que no faré. Primer: no us diré si la IA és bona o dolenta. Hi ha arguments legítims en contra — el consum energètic, la qüestió dels drets d'autor de les dades d'entrenament, el desplaçament laboral. Són preocupacions reals i no les ignoraré, però avui no és el debat.

Segon: no us donaré una llista de 50 eines amb noms que canviaran d'aquí a tres mesos. El que sí que faré és proposar-vos un marc mental. Perquè les eines canvien cada setmana, però un bon marc per pensar-hi dura molt més."
-->

---

# Preocupacions legítimes

El consum energètic és real.
La qüestió dels drets d'autor no està resolta.
El desplaçament laboral és una possibilitat concreta.

Cap d'aquestes és una raó per no mirar-hi de cara.

<!--
[continuació disclaimers — 30 segons]

"Reconèixer els problemes no és el mateix que fugir-ne. Com veurem, la història ens ensenya que la gent que millor ha navegat les transformacions tecnològiques no és la que s'hi ha oposat ni la que l'ha abraçada cegament — és la que ha preguntat: a qui beneficia això, i com podem fer que ens beneficiï a nosaltres?"
-->

---

# Catalunya, 1790

Un fuster de Berga inventa la **berguedana**.

130 fusos. Feta de fusta de pi.
Operada a mà.

La *spinning jenny* catalana.

<!--
[CONTEXT HISTÒRIC — inici, 1 minut]

"Posem-nos en context. Catalunya, finals del segle XVIII. Ramon Farguell, un fuster de Berga, construeix la berguedana — una màquina de filar de fusta de pi amb 130 fusos. És una adaptació millorada de la spinning jenny anglesa. Centenars de tallers catalans l'adopten. La filatura local explota.

Però la berguedana només era el primer pas."
-->

---

# 1832: El Vapor Bonaplata

Primera fàbrica a vapor de la Península Ibèrica.
Barcelona, carrer dels Tallers.
600-700 treballadors.

1835: cremada durant les bullangues.

<!--
[CONTEXT HISTÒRIC — Bonaplata, 1 minut]

"El 1832, al carrer dels Tallers de Barcelona, s'inaugura el Vapor Bonaplata. Primera fàbrica moguda per vapor de tota la Península. 600 treballadors, filatura, teixidura i foneria — tot integrat.

Tres anys després, la nit del 5 d'agost de 1835, durant les bullangues, la cremen. Un mariner, Narcís Pardinas, és afusellat dos dies més tard. La família Bonaplata marxa de Barcelona per no tornar-hi mai."
-->

---

# 1844: Arriba la selfactina

Completament automatitzada.
Fins a 1.000 fusos.
Un sol operari.

Substitueix filadors adults per dones i nens
a sous molt inferiors.

<!--
[CONTEXT HISTÒRIC — selfactina, 1 minut]

"El 1844 arriba la selfactina — la self-acting mule. Completament automatitzada, fins a mil fusos, un sol operari. I aquí és on la cosa canvia de veritat: la selfactina permetia substituir filadors adults qualificats per dones i nens, a sous molt inferiors.

En cinc anys, els fusos selfactins passen de 91.000 a més de 200.000. La productivitat puja, els salaris s'enfonsen."
-->

---

# 1854: La selfactinada

Els obrers cremen fàbriques amb selfactines.

El capità general **prohibeix** les selfactines.

Madrid revoca la prohibició...
però no la publica fins **8 mesos després**.
Per por.

<!--
[CONTEXT HISTÒRIC — selfactinada, 1 minut]

"El juliol de 1854, els obrers cremen fàbriques que tenen selfactines. La resposta és brutal — tres persones afusellades — però el capità general, per calmar la revolta, fa una cosa extraordinària: prohibeix les selfactines.

Madrid revoca la prohibició en dues setmanes... però no la publica fins al maig de 1855. Vuit mesos de silenci. Per por de la reacció obrera.

Fixeu-vos: la tecnologia no es va aturar. Però la revolta va forçar una negociació."
-->

---

# El que demanaven els ludites

No demanaven destruir les màquines.

Demanaven:

**Qui s'endú el benefici?**
**Qui paga el cost?**
**On és la nostra dignitat?**

<!--
[LLIÇÓ LUDITA — 2 minuts]

"Ara, la pregunta clau. Quan sentim 'ludites', pensem en gent trencant màquines. Però les reclamacions dels ludites catalans — i dels anglesos — no anaven sobre les màquines. Anaven sobre tres preguntes:

Qui s'endú el benefici d'aquesta productivitat? Qui paga el cost de la transició? I on queda la dignitat del treballador?

El conflicte de les selfactines va acabar amb el reconeixement oficial de les Societats Obreres. El 1855, cent mil treballadors fan la primera vaga general de la història d'Espanya sota el lema 'Associació o mort'. No demanaven tornar enrere — demanaven tenir veu en com anava endavant.

I aquí ve el clímax: el 1919, la vaga de la Canadenca. Comença per 8 oficinistes als quals els retallen el sou. Acaba amb Barcelona a les fosques, sense tramvies, sense diaris. I el resultat? El 3 d'abril de 1919, Espanya es converteix en el primer país d'Europa occidental a legislar la jornada de vuit hores. Vint dies abans que França.

De la berguedana de fusta de 1790 a la jornada de vuit hores de 1919: 129 anys de lluita. I les tres preguntes dels ludites mai es van respondre soles. Es van respondre perquè algú les va fer."
-->

---

# Tres preguntes, dos segles després

**Qui s'endú el benefici?**
**Qui paga el cost?**
**On és la nostra dignitat?**

<!--
[TRANSICIÓ AL PRESENT — 30 segons]

"Tornem al 2026. Les tres preguntes són exactament les mateixes.

Qui s'endú el benefici de la productivitat dels LLMs? Qui paga el cost de la transició? I on queda la dignitat de l'enginyer?

Anem a veure-ho."
-->

---

# El paral·lelisme

| Tèxtil s. XIX | Programari s. XXI |
|---|---|
| Berguedana | Autocomplete, Copilot |
| Selfactina | Agents de codi |
| Filador manual | Programador "escriptor de codi" |
| Teixidor de sistemes | Enginyer de programari |

<!--
[PARAL·LELISME — 2 minuts]

"La berguedana era una eina manual millorada. Com l'autocomplete o el primer Copilot: t'ajudava, però tu feies la feina. La selfactina era completament automàtica. Com els agents de codi actuals: tu dones la instrucció, la màquina executa.

I aquí és on el paral·lel es posa interessant. Quan la spinning jenny va automatitzar la filatura, els teixidors manuals van viure 20 anys d'or. La seva feina s'havia convertit en el coll d'ampolla — tothom filava ràpid, però algú havia de teixir tot aquell fil. Fins que van automatitzar la teixidura.

Avui, el codi es fila sol. El coll d'ampolla és el disseny de sistemes, la revisió, el judici. Això és teixir. La pregunta és: som al principi dels 20 anys d'or per a teixidors de sistemes?"
-->

---

# El codi es fila sol

Produir codi és gairebé gratuït.

Produir **bon** codi segueix sent car.

<!--
[BON CODI — 1.5 minuts]

"Deixeu-me ser clar: produir codi mai havia estat tan barat. Però produir bon codi — codi que funciona, que és segur, que gestiona els errors, que té tests, que es pot mantenir — això segueix sent car.

Un agent et genera un PR de 500 línies en 10 minuts. Però algú ha de revisar aquelles 500 línies. Algú ha de decidir si solucionen el problema correcte. Algú ha d'entendre si aquella correcció de seguretat realment tanca el vector d'atac o simplement ho sembla.

El primer 90% del codi representa el primer 90% del temps de desenvolupament. L'altre 10% del codi representa l'altre 90% del temps. Ara la primera part és gairebé gratis. L'altra segueix igual."
-->

---

# Escriure codi ≠ enginyeria de programari

Els agents escriuen codi.

No fan enginyeria de programari.

L'enginyeria és *trade-offs*, context,
decisions que cap model té.

<!--
[DISTINCIÓ CLAU — 1.5 minuts]

"Hi ha una distinció fonamental que hem de tenir clara: escriure codi i fer enginyeria de programari no és la mateixa feina.

Escriure codi és filatura. L'enginyeria de programari és teixidura. És decidir trade-offs. És entendre el domini del negoci, l'estratègia de producte, el deute tècnic. És la conversa que l'equip va tenir la setmana passada sobre per què vam triar una arquitectura i no una altra.

El desenvolupament de programari tal com el coneixíem és mort. Però l'enginyeria de programari està més viva que mai. El rol s'ha transformat: d'escriure codi a dissenyar sistemes que escriuen codi."
-->

---

# Per què els LLMs són aquí per quedar-se

No es tracta del model.
Es tracta del **procés**.

Un bon model amb mal procés perd contra
un model decent amb bon procés.

Sempre.

<!--
[LLMs PER QUEDAR-SE — 2 minuts]

"I els LLMs no marxaran. Però no per la raó que penseu.

No es tracta del model. Es tracta del procés. Un bon model amb mal procés perd contra un model decent amb bon procés. Sempre. Igual que en el desenvolupament tradicional: no tots els programadors són excepcionals, però un equip de programadors competents amb un bon procés pot construir un gran producte.

El resultat el defineix el procés, no el model. El model és una peça del puzle.

I hi ha una cosa més: no tot han de ser LLMs al núvol."
-->

---

# No tot és el núvol

**NPUs locals**: processadors neuronals als xips actuals

**Clústers locals**: 4 Mac Studios per 12.000-50.000 €
vs. 100.000+ € en servidors GPU.
Cap cost recurrent. Cap cost per token.

**Models oberts**: cada cop més capaços

<!--
[ALTERNATIVES LOCALS — 1.5 minuts]

"Quan parlem d'IA generativa, la majoria pensem en APIs al núvol. Però el panorama està canviant. Els processadors moderns ja porten NPUs — unitats de processament neuronal integrades. Hi ha projectes com Taalas que aprofiten aquestes NPUs per executar models localment.

I els clústers locals estan sent cada cop més viables. Un clúster de 4 Mac Studios costa entre 12.000 i 50.000 euros, versus 100.000 o més en servidors GPU equivalents. És una despesa de capital única — sense costos recurrents, sense costos per token.

Això importa perquè canvia la resposta a la primera pregunta ludita: 'qui s'endú el benefici?' Si el model corre a la teva màquina, el benefici és teu."
-->

---

# Com evolucionar

<!--
[TRANSICIÓ — 10 segons]

"Aleshores, com naveguem tot això? Us proposo quatre principis."
-->

---

# 1. Judici sobre velocitat

La capacitat que val més ja no és escriure ràpid.

És **saber si el que s'ha escrit és correcte**.

Llegiu codi. Molt codi.
Codi d'altres, codi generat, codi open source.

<!--
[PRINCIPI 1 — 1.5 minuts]

"Primer: judici sobre velocitat. La meva capacitat de lliurar ja no està limitada per la velocitat a la que codifico. Està limitada per la meva habilitat per revisar. I crec que és exactament com ha de ser.

Quan corregeixo una vulnerabilitat de seguretat, no miro només si els tests passen. Em pregunto: realment tanca el vector d'atac? Hi ha casos límit que l'agent ha passat per alt? És la correcció correcta, o simplement una correcció?

La habilitat més valuosa del 2026 és llegir codi. Molt codi. Codi open source, codi de producció, codi generat per IA. Entendre què fa, què no fa, i què fa malament."
-->

---

# 2. Processos, no eines

Escriviu plans abans que l'agent escrigui codi.
PRs petits amb descripcions clares.
Verificació automatitzada.

Si corregiu el mateix error tres cops,
escriviu un test.

<!--
[PRINCIPI 2 — 1.5 minuts]

"Segon: invertiu en processos, no en eines. Les eines canvien cada trimestre. Un bon procés dura anys.

Algunes coses concretes: no deixeu que l'agent escrigui codi fins que hàgiu revisat i aprovat un pla escrit. PRs petits amb descripcions clares — és l'única combinació que passa la revisió a un ritme raonable. Si tres revisors de dos equips troben el mateix tipus d'error, això és un check esperant a ser escrit.

I sobretot: no envieu PRs amb codi que no heu revisat vosaltres. Si obriu un PR de 500 línies generades per un agent i no heu fet la feina d'assegurar-vos que funciona, esteu delegant la feina real als vostres companys."
-->

---

# 3. Pregunteu-vos a qui beneficia

L'eina **amplifica** o **substitueix**?
Profunditza el vostre pensament o
genera output que accepteu tal qual?

Les mateixes tres preguntes. Sempre.

<!--
[PRINCIPI 3 — 1.5 minuts]

"Tercer: cada cop que adopteu una eina, pregunteu-vos a qui beneficia.

Hi ha dues maneres d'aplicar la IA al treball. Una: generar l'output. Escrius un prompt, l'eina produeix un esborrany, tu passes d'autor a editor de producció aliena. Ràpid, eficient. Però no has modelat el treball. L'has acceptat.

L'altra: aprofundir el procés. L'eina et fa preguntes que treuen el que realment penses. Et mostra tres angles que no havies considerat. Desafia l'assumpció més feble. T'ajuda a anar més endins — i el resultat és genuïnament teu.

La mateixa tecnologia, dues aplicacions oposades. La diferència és si l'eina tracta la participació humana com el coll d'ampolla o com el propòsit."
-->

---

# 4. Responsabilitat

Un ordinador no pot ser mai responsable.

*— IBM, formació interna, 1979*

Quan lliuro codi, el meu nom hi és.
Puc usar IA per anar més ràpid.
No puc externalitzar el meu judici.

<!--
[PRINCIPI 4 — 1 minut]

"I quart: la responsabilitat és vostra i no es delega.

IBM ho va dir el 1979 en una formació interna: 'Un ordinador no pot ser mai responsable.' Mai ha estat més rellevant.

Quan lliuro codi, el meu nom hi és. Quan hi ha una vulnerabilitat a Undici o un bug a Fastify, és la meva responsabilitat. Puc usar IA per anar més ràpid, però no puc externalitzar el meu judici. No puc externalitzar la meva responsabilitat."
-->

---

# Recapitulant

**Qui s'endú el benefici?** Pregunteu-ho sempre.

**Com afrontar la nova realitat?** Judici, processos,
pensament crític. Teixidors, no filadors.

**No tot és el núvol.** NPUs, clústers locals,
models oberts. El benefici pot ser vostre.

<!--
[RECAPITULACIÓ — 1 minut]

"Tres idees per endur-vos.

Primera: pregunteu-vos sempre qui s'endú el benefici. Exactament com ho feien els ludites. No és una pregunta retrògrada — és la pregunta que va acabar guanyant la jornada de vuit hores.

Segona: el futur és per als teixidors de sistemes. Judici sobre velocitat, processos sobre eines, pensament crític sobre productivitat cega. Llegiu codi, reviseu codi, enteneu codi.

Tercera: no tot ha de córrer al núvol d'algú altre. Les NPUs, els clústers locals, els models oberts estan aquí. Quan el model corre a casa teva, la resposta a 'qui s'endú el benefici?' ets tu."
-->

---

# De filadors de codi
# a teixidors de sistemes

La pregunta no és si les màquines filaran per nosaltres.

La pregunta és **qui teixeix**.

<!--
[COMIAT — 1 minut]

"La revolució industrial tèxtil a Catalunya va ser molt més que un canvi tecnològic. Va ser una fractura. Les màquines van crear riquesa sense precedents, però el cost humà va ser immens. I la resistència obrera va evolucionar: de cremar màquines a organitzar-se, de l'acció directa a la negociació col·lectiva.

Nosaltres estem vivint la revolució industrial del programari. El codi ja es fila sol. La pregunta no és si les màquines filaran per nosaltres — ja ho fan. La pregunta és qui teixeix. I qui s'endú el benefici del que es teixeix.

Us ho van dir fa 200 anys des dels carrers de Barcelona: Associació o mort. Jo ho traduiria per al 2026: Criteri o irrellevància.

Gràcies."
-->

---

<p class="small">

Dario Castañé — @darccio

</p>

<!--
Diapositiva final amb contacte. No cal dir res, només deixar-la visible durant les preguntes.
-->
