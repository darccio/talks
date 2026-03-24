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

- +700$ en tokens de Claude
- Productivitat x2 (o x3, podcasts)
- Plantejar-me que vol dir ser enginyer
-->

---

## Dario Castañé

Enginyer de programari sènior @ Datadog

<!--
[PRESENTACIÓ — 20 segons]

- Datadog
- Enginyer de programari en llibreries i eines de codi lliure
- Principalment en Go (15 anys d'usuari, 8 anys professionalment)
- Avui vull explicar-vos com plantejo el futur professional
-->

---

# Avui proposaré un marc mental.

No una llista d'eines. No un judici moral.

<!--
[DISCLAIMERS — 2 minuts]

No faré...

Primer:
- No us diré si la IA generativa és bona o dolenta
- Hi ha arguments legítims en contra: consum, drets d'autor, impacte en el mercat laboral
- Preocupacions reals però no el focus de la meva xerrada

Segon:
- No llista de 50 eines que canviaran aviat
- Però després al torn de preguntes em podeu preguntar

Sí faré:
- Proposar un marc mental
- Inspirat en els resultats de la lluita obrera del sector tèxtil entre els segles XVIII i XX

Fem un breu repàs del que va succeïr...
-->

---

# De 130 fusos a 1.000

En menys de 70 anys, la filatura passa
de manual a completament automàtica.

<!--
[CONTEXT HISTÒRIC — 2 minuts]

- Abans de la Revolució Industrial...
- Finals del S. XVIII
- Berguedana, màquina de filar (converteix cotó a fil)
- Tot i ser manual, permet passar d'un fus a 130 fusos
- 70 anys després, en plena Revolució Industrial...
- Selfactina, automatització, 1.000 fusos, un sol operari
- Productivitat puja, salaris s'enfonsen

La conseqüència?
-->

---

# 1854: cremen les fàbriques.

Es prohibeixen les selfactines.
Madrid deroga la prohibició dues setmanes després.
No la publica fins **8 mesos més tard**.
Per por a la revolta obrera.

Però aquesta no va impedir que la tecnologia s'implantés.

<!--
[CONTEXT HISTÒRIC — selfactinada, 1 minut]

- Màquines cremant
- Selfactines prohibides, però legalitzades en no res
- La tecnologia no es va aturar
- Tots vestim roba teixida per màquines

I si us pregunto per algú que va destacar per aquesta lluita, quin nom us surt?
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

- Quan escoltem aquest nom, pensem en gent destruint màquines
- Però les seves reclamacions no eren sobre aquestes
- Sinó sobre aquestes tres preguntes
- Qui se'n beneficiava? Qui en pagava el cost? I on era la seva dignitat?

- Aquests conflictes van obtenir victòries
- Selfactinada: reconeixement de societats obreres, vaga general, associació o mort
- Vaga de la Canadenca: jornada de vuit hores

I que té a veure tot aix`o amb la IA generativa?
-->

---

# Les mateixes preguntes, dos segles després

**Qui s'endú el benefici?**
**Qui paga el cost?**
**On queda la nostra professió?**

<!--
[TRANSICIÓ AL PRESENT — 30 segons]

- 2026: les tres preguntes segueixen en vigor
- No buscarem la resposta a les tres
- Però sí que hem de veure com el que va passar al tèxtil català és un mirall del nostre present
- I que som al mateix punt però no ens hem de desanimar
-->

---

# El paral·lelisme

Berguedana → Autocompletar, Copilot
Selfactina → Agents de codi

Filador manual → Programador "escriptor de codi"
Teixidor → Enginyer de programari

<!--
[PARAL·LELISME — 2 minuts]

- Berguedana, eina manual millorada, autocompletar o el primer Copilot: ajuda però tu fas la feina
- Selfactina, completament automàtica, com els agents de codi: implementen el que els hi demanes
- El paral·lel es posa interessant...
- Amb la berguedana, 20 anys d'or dels teixidors
- Coll d'ampolla fins que es va automatitzar la teixidura

- Avui és el codi el que es fila sol
- Coll d'ampolla: disseny de sistemes a escola, la revisió, la capacitat de jutjar
- Som a l'inici de dues dècades daurades per a teixidors de sistemes?
-->

---

# El codi es fila sol

Si el codi és molt més barat de produir,
**què és el que té valor?**

<!--
[BON CODI — 1.5 minuts]

- Produir codi mai ha estat tan barat
- Què te valor? La capacitat de jutjar
- Jutjar que el codi generat realment...
- Soluciona el problema, és segur, és mantenible...
- I això és el que és encara car.

- PR de 500 línies en minuts, però algú l'ha de revisar
- Perquè escriure codi no és enginyeria de programari
-->

---

# Escriure codi ≠ enginyeria de programari

Vibe coding: programar per intuició, acceptant el resultat.
Agentic engineering (o automatic programming): dissenyar sistemes i revisar el resultat abans d'integrar-lo.

Els agents escriuen codi.
No decideixen trade-offs.
Necessiten guardaraïls estrictes.

<!--
[DISTINCIÓ CLAU — 2 minuts]

- En aquest nou món tenim...
- Vibe coding: et deixes portar, genial per a prototips, però no és enginyeria

- L'enginyeria es troba en el següent nivell, l'agentic engineering
- Decidir compromisos, dissenyar amb patrons
- Entendre el domini del negoci, l'estratègia de producte, el deute tècnic...
- En això els humans encara guanyem perquè tenim agència

- El desenvolupament de programari tal com el coneixíem és mort
- L'enginyeria de programari està més viva que mai
- El rol s'ha transformat: d'escriure codi a dissenyar sistemes des de patrons, amb codi generat automàticament
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
