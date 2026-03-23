---
marp: true
theme: default
paginate: true
style: |
  section {
    background-color: #1E1E1E !important;
    color: #F5F5F5 !important;
    font-family: 'Helvetica Neue', Arial, sans-serif;
    font-size: 36px;
  }
  h1, h2 {
    color: #FFFFFF !important;
    font-size: 52px;
  }
  th, td {
    background-color: #2A2A2A !important;
    color: #F5F5F5 !important;
    border-color: #555 !important;
  }
  section.invert {
    background-color: #F5F5F5 !important;
    color: #1E1E1E !important;
  }
  section.invert h1,
  section.invert h2 {
    color: #111111 !important;
  }
  section.invert th,
  section.invert td {
    background-color: #E8E8E8 !important;
    color: #1E1E1E !important;
    border-color: #AAA !important;
  }
---

<!-- _paginate: false -->

# De filadors de codi a teixidors de sistemes

La revolució industrial del programari

<!--
[EXORDIO — 1 minut]

Començar directe, sense salutació llarga.

"No us menteixo si us dic que si m'haguessin dit fa un any que en aquest últim mes hauria gastat a la feina 700$ en tokens, no m'ho hauria cregut.

Però així ha estat i puc dir que he multiplicat la meva productivitat per dos. Ara puc fer dues coses alhora... mentre espero que un agent acabi de fer el que li he demanat. Tres si comptem escoltar podcasts mentre treballo.

Un cop he arribat a aquest punt m'he vist obligat a repensar què significa ser enginyer de programari avui."
-->

---

## Dario Castañé

Enginyer de programari sènior @ Datadog

<!--
[PRESENTACIÓ — 30 segons]

"Em dic Dario Castañé. Treballo a Datadog com a enginyer de programari, concretament en projectes de codi obert. Mantinc dd-trace-go i Orchestrion, dues eines d'instrumentació en Go, i també mantinc mergo, una llibreria Go amb 3.000 estrelles i més de 30.000 repositoris utilitzant-la.

I aquest mes, personalment estic treballant en un projecte en Rust. Sóc gopher de fa vora 15 anys, 8 anys usant-lo a diari professionalment. Ni fava de Rust. N'aprenc preguntant a l'agent mentre desenvolupo una aplicació d'escriptori. Totalment fora del meu element. Però amb agents puc fer el que abans m'hagués requerit tones d'esforç."
-->

---

# Avui proposaré un marc mental.

No una llista d'eines. No un judici moral.

<!--
[DISCLAIMERS — 2 minuts]

"Abans de res, dues coses que no faré. Primer: no us diré si la IA és bona o dolenta. Hi ha arguments legítims en contra — el consum energètic, la qüestió dels drets d'autor de les dades d'entrenament, el desplaçament laboral. Són preocupacions reals i no les ignoraré, però avui no és el debat.

Segon: no us donaré una llista de 50 eines amb noms que canviaran d'aquí a tres mesos. El que sí que faré és proposar-vos un marc mental. Perquè les eines canvien cada setmana, però un bon marc per pensar-hi dura molt més."

Com veurem, la història ens ensenya que la gent que millor ha navegat les transformacions tecnològiques no és la que s'hi ha oposat ni la que l'ha abraçada cegament — és la que ha preguntat: a qui beneficia això, i com podem fer que ens beneficiï col·lectivament?
-->


---

# 1790 → 1844

Berguedana: 130 fusos, operada a mà.
Selfactina: 1.000 fusos, un sol operari.

<!--
[CONTEXT HISTÒRIC — 2 minuts]

"Posem-nos en context. Ramon Farguell, un fuster de Berga, construeix la berguedana — una adaptació millorada de la spinning jenny anglesa, una màquina de filar, que converteix cotó en fil per teixir. Centenars de tallers catalans l'adopten perquè a mà només pots fer un fus, mentre que aquesta en permetia 130 alhora.

El 1844 arriba la selfactina. Completament automatitzada, fins a mil fusos, un sol operari. Cinquanta-quatre anys, de 130 fusos manuals a 1.000 d'automatitzats. La productivitat puja. Els salaris s'enfonsen."
-->

---

# 1854: cremen les fàbriques.

Es prohibeixen les selfactines.
Madrid deroga la prohibició dues setmanes més tard.
No la publica fins **8 mesos després**.
Per por.

<!--
[CONTEXT HISTÒRIC — selfactinada, 1 minut]

"El juliol de 1854, els obrers cremen fàbriques que tenen selfactines. La resposta és brutal — tres persones afusellades — però el poder, per calmar la revolta, fa una cosa extraordinària: prohibeix les selfactines.

Madrid revoca la prohibició en dues setmanes... però no la publica fins al maig de 1855. Vuit mesos de silenci. Per por de la reacció obrera.

Fixeu-vos: la tecnologia no es va aturar. Tots aquí vestim roba teixida per màquines. Però la revolta va forçar una negociació.

Potser alguns de vosaltres ja teniu al cap un nom: els ludites."
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

"Quan sentim 'ludites', pensem en gent trencant màquines. Però les reclamacions dels ludites catalans — i d'altres indrets — no anaven sobre les màquines. Anaven sobre tres preguntes:

Qui s'endú el benefici d'aquesta productivitat? Qui paga el cost de la transició? I on queda la dignitat del treballador?

El conflicte de les selfactines va acabar amb el reconeixement oficial de les Societats Obreres, la llavor dels sindicats moderns, amb una vaga general sota el lema "Associació o mort".

I més tard el 1919, la vaga de la Canadenca acaba per portar la regulació de la jornada de vuit hores.

De la berguedana de fusta de 1790 a la jornada de vuit hores de 1919: 129 anys de lluita. I les tres preguntes dels ludites mai es van respondre soles. Es van respondre perquè algú les va fer."
-->

---

# Les mateixes preguntes, dos segles després

**Qui s'endú el benefici?**
**Qui paga el cost?**
**On queda la nostra professió?**

<!--
[TRANSICIÓ AL PRESENT — 30 segons]

"Tornem al 2026. Les tres preguntes són exactament les mateixes.

Qui s'endú el benefici de la productivitat dels LLMs? Qui paga el cost de la transició? I on queda la professió d'enginyer?

Anem a veure-ho."
-->

---

# El paral·lelisme

| Tèxtil s. XVIII-XIX | Programari s. XXI |
|---|---|
| Berguedana | Autocomplete, Copilot |
| Selfactina | Agents de codi |
| Filador manual | Programador "escriptor de codi" |
| Teixidor | Enginyer de programari |

<!--
[PARAL·LELISME — 2 minuts]

"La berguedana era una eina manual millorada del s. XVIII. Com l'autocomplete o el primer Copilot: t'ajudava, però tu feies la feina. La selfactina - del s. XIX - era completament automàtica. Com els agents de codi actuals: tu dones la instrucció, la màquina executa.

I aquí és on el paral·lel es posa interessant. Quan la berguedana va automatitzar la filatura, els teixidors manuals van viure 20 anys d'or. La seva feina s'havia convertit en el coll d'ampolla — tothom filava ràpid, però algú havia de teixir tot aquell fil. Fins que van automatitzar la teixidura.

Avui, el codi es fila sol. El coll d'ampolla és el disseny de sistemes, la revisió, el judici. Això és teixir. La pregunta és: som al principi dels 20 anys d'or per a teixidors de sistemes?"
-->

---

# El codi es fila sol

Produir codi ha esdevingut gairebé gratuït.

Produir **bon** codi segueix sent car.

<!--
[BON CODI — 1.5 minuts]

"Deixeu-me ser clar: produir codi mai havia estat tan barat. Però produir bon codi — codi que funciona, que és segur, que gestiona els errors, que té tests, que es pot mantenir — això segueix sent car.

Un agent et genera un PR de 500 línies en 10 minuts. Però algú ha de revisar aquelles 500 línies. Algú ha de decidir si solucionen el problema correcte. Algú ha d'entendre si aquella correcció de seguretat realment tanca el vector d'atac o simplement ho sembla.

Perquè escriure codi no és enginyeria de programari."
-->

---

# Escriure codi ≠ enginyeria de programari

Els agents escriuen codi.
No decideixen trade-offs.
Necessiten guardaraïls estrictes.

<!--
[DISTINCIÓ CLAU — 1.5 minuts]

"Hi ha una distinció fonamental que hem de tenir clara: escriure codi i fer enginyeria de programari no és la mateixa feina.

Escriure codi és filatura. L'enginyeria de programari és teixidura. És decidir compromisos. És entendre el domini del negoci, l'estratègia de producte, el deute tècnic. És la conversa que l'equip va tenir la setmana passada sobre per què vam triar una arquitectura i no una altra.

El desenvolupament de programari tal com el coneixíem és mort. Però l'enginyeria de programari està més viva que mai. El rol s'ha transformat: d'escriure codi a dissenyar sistemes que escriuen codi."
-->

---

# Agents ≠ enginyeria de programari

Vibe coding: programar per intuició.
Agentic engineering.
Automatic programming.

<!--
TODO
-->

---

# No es tracta del model.

Es tracta del **procés**.

<!--
[LLMs PER QUEDAR-SE — 2 minuts]

"I els LLMs no marxaran. Però no per la raó que penseu.

No es tracta del model. Es tracta del procés. Un bon model amb mal procés perd contra un model decent amb bon procés. Sempre. Igual que en el desenvolupament tradicional: no tots els programadors són excepcionals, però un equip de programadors competents amb un bon procés pot construir un gran producte.

El resultat el defineix el procés, no el model. El model és una peça del puzle.

I hi ha una cosa més: no tot han de ser LLMs al núvol."
-->

---

# Maquinari local

NPUs.
Clústers de 4 Mac Studios: 12.000–50.000 €.
Dispositius d'inferència local.

<!--
[ALTERNATIVES LOCALS 1/2 — 45 segons]

"Els processadors moderns ja porten NPUs — unitats de processament neuronal integrades. I els clústers locals estan sent cada cop més viables: 4 Mac Studios costen entre 12.000 i 50.000 euros, versus 100.000 o més en servidors GPU equivalents. Sense costos recurrents, sense costos per token."

També hi ha dispositius com Tiiny - actualment en Kickstarter - o les Taalas hardcore model que s'estan posicionant per inferència local.
-->

---

# Models oberts

Cada cop més capaços. Cada cop més accessibles.

<!--
[ALTERNATIVES LOCALS 2/2 — 30 segons]

"Això importa perquè canvia la resposta a la primera pregunta ludita: 'qui s'endú el benefici?' Si el model corre a la teva màquina, el benefici és teu."
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

<!--
[PRINCIPI 1 — 1.5 minuts]

"Primer: judici sobre velocitat. La meva capacitat de lliurar ja no està limitada per la velocitat a la que codifico. Està limitada per la meva habilitat per revisar. I crec que és exactament com ha de ser.

Quan corregeixo un bug, no miro només si els tests passen. Em pregunto: realment resol la causa arrel? Hi ha casos límit que l'agent ha passat per alt? És la correcció correcta, o simplement un pedaç?

L'habilitat més valuosa del 2026 és llegir codi. Molt codi. Codi open source, codi de producció, codi generat per IA. Entendre què fa, què no fa, i què fa malament."
-->

---

# 2. Processos, no eines

Les eines canvien cada trimestre.
Un bon procés dura anys.

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

*Un ordinador no pot ser mai responsable.*

— IBM, formació interna, 1979

<!--
[PRINCIPI 4 — 1 minut]

"I quart: la responsabilitat és vostra i no es delega.

IBM ho va dir el 1979 en una formació interna: 'Un ordinador no pot ser mai responsable.' Mai ha estat més rellevant.

Quan lliuro codi, el meu nom hi és. Quan hi ha una vulnerabilitat a Undici o un bug a Fastify, és la meva responsabilitat. Puc usar IA per anar més ràpid, però no puc externalitzar el meu judici. No puc externalitzar la meva responsabilitat."
-->

---

# Criteri o irrellevància.

<!--
[RECAPITULACIÓ — 1 minut]

"Tres idees per endur-vos.

Primera: pregunteu-vos sempre qui s'endú el benefici. Exactament com ho feien els ludites. No és una pregunta retrògrada — és la pregunta que va acabar guanyant la jornada de vuit hores.

Segona: el futur és per als teixidors de sistemes. Judici sobre velocitat, processos sobre eines, pensament crític sobre productivitat cega. Llegiu codi, reviseu codi, enteneu codi.

Tercera: no tot ha de córrer al núvol d'algú altre. Les NPUs, els clústers locals, els models oberts estan aquí. Quan el model corre a casa teva, la resposta a 'qui s'endú el benefici?' ets tu."
-->

---

<!-- _paginate: false -->

# De filadors de codi<br>a teixidors de sistemes

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

<!-- _paginate: false -->

<p class="small">

Dario Castañé — @darccio

cc@dario.cat
dario.cat

</p>

<!--
Diapositiva final amb contacte. No cal dir res, només deixar-la visible durant les preguntes.
-->
