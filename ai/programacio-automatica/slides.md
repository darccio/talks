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

# Avui proposaré un marc mental

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

# 1854: cremen les fàbriques

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
No decideixen compromisos.
Necessiten guardaraïls estrictes.
No operen en el buit

<!--
[DISTINCIÓ CLAU — 2 minuts]

- En aquest nou món tenim...
- Vibe coding: et deixes portar, genial per a prototips, però no és enginyeria

- L'enginyeria es troba en el següent nivell, l'agentic engineering
- Decidir compromisos, dissenyar amb patrons
- Entendre el domini del negoci, l'estratègia de producte, el deute tècnic, l'excel·lència operacional...
- També detectar quan és necessari posar un guardaraïls: linters, formatadors, ...
- Claude té hooks que executa sempre, Git també
- En això els humans encara guanyem perquè tenim agència

- El desenvolupament de programari tal com el coneixíem és mort
- L'enginyeria de programari està més viva que mai
- El rol s'ha transformat: d'escriure codi a dissenyar sistemes des de patrons, amb codi generat automàticament
- Però també cal desplegar i garantir que aquests programaris funcionen (SRE com perfil de futur)
-->

---

# L'enginyeria canvia, les eines també

Però aquestes eines són al núvol.
Serem captius?

No necessàriament.

<!--
[TRANSICIÓ - 1 minut]

- I els LLMs són aquí per quedar-se
- Però en quines condicions?
- No es pot negar que la IA generativa ara viu al núvol
- Però s'albiren desenvolupament que permetin una IA generativa local
- Sostenible en el seu ús (l'entrenament és més difícil)
- Com és possible?
-->

---

# Maquinari local i models oberts

NPUs integrades als processadors moderns.
Clústers locals cada cop més viables.
Models oberts cada cop més capaços.

Si el model corre a les nostres màquines,
**el benefici és nostre**.

<!--
- Apple és capdavantera en inferència local
- NPUs: Neural Engine
- RDMA sobre Thunderbolt: Remote Direct Memory Access
- Ja es poden muntar clústers de Mac Studios, per una quarta part del que pot costar fer l'equivalent amb GPUs i servidors
- Però hi ha vida més enllà de la poma
- Tiiny (Kickstarter, models fins 120 milers de milions de paràmetres)
- Taalas Hardcore model
- No estem comdemnats a passar per caixa eternament (pagar-ne el cost)

- A banda, hi ha models oberts que poden funcionar en inferència local
- Models oberts prou potents per la majoria
- No tothom necessita utilitzar models de frontera
- Perquè no tot es tracta d'on còrre el model o el model en si
-->

---

# Però no es tracta del model o on s'executa aquest

Es tracta del **procés**.

<!--
[LLMs PER QUEDAR-SE — 1.5 minuts]

- Es tracta del procés
- I el millor ara mateix és començar a definir el procés que ens funcioni
- Un bon model amb un mal procés perdrà sempre contra un model decent amb bon procés
- Això és també cert pel desenvolupament tradicional
- Un equip competent amb bon procés és millor que un equip excepcional caòtic
- En resum, el resultat el defineix el procés, no el model
- El model és una peça del puzle.
- Parlem dels principis que considero que ha de tenir un bon procés d'enginyeria
-->

---

# 1. Capacitat de jutjar, no velocitat

La capacitat que val més ja no és escriure ràpid.

És **saber si el que s'ha escrit és correcte**.

<!--
[PRINCIPI 1 — 1.5 minuts]

Primer:
- Saber jutjar per sobre de la velocitat
- No més coll d'ampolla en programar
- Coll d'ampolla en revisar
- Així ha de ser
- Exemple personal: corb
- L'habilitat més valuosa: llegir codi (lliure, producció, generat per IA)
- Entendre què fa, què no fa, i què fa malament
-->

---

# 2. Bones pràctiques, no eines

Les eines canvien cada trimestre.
Les bones pràctiques duren anys.

<!--
[PRINCIPI 2 — 1.5 minuts]

Segon:
- Invertir en bones pràctiques, no eines
- Les eines canvien constantment ara mateix
- Idees:
- No generar codi fins haver revisat i aprovat un pla exhaustiu
- PRs petites amb descripcions clares: demostrat com la combinació perfecta que passa revisió a un ritme raonable
- Mai enviar PRs generades sense revisar-ne el codi
- Crear guardaraïls personalitzats: linters
- Automatitzar-los en les pipelines de CI i en local
-->

---

# 3. Pregunteu-vos a qui beneficia

L'eina **amplifica** o **substitueix**?
Profunditza el pensament o
genera resultats que acceptem tal qual?

<!--
[PRINCIPI 3 — 1.5 minuts]

Tercer:
- Quan adoptem una eina, preguntar-se a qui beneficia
- Tant econòmicament com professionalment (sobretot això, ens ajuda a crèixer?)
- Dues maneres de treballar amb la IA:
1. Generar el resultat: prompt, resultat, acceptació humana passiva
2. Aprofundir en el procés: l'eina et fa preguntes, et mostra diferents angles, et desafia les assumpcions
- La mateixa tecnologia, dues aplicacions oposades: participació humana com coll d'ampolla o propòsit central"
-->

---

# 4. Responsabilitat

*Un ordinador no pot ser mai responsable.*

— IBM, formació interna, 1979

<!--
[PRINCIPI 4 — 1 minut]

Quart:
- La responsabilitat és nostra i no es delega
- IBM, formació interna, 1979: 'Un ordinador no pot ser mai responsable'
- Quan lliuro codi, el meu nom hi és, la responsabilitat és meva i no es pot externalitzat
-->

---

# Criteri o irrellevància.

<!--
[RECAPITULACIÓ — 1 minut]

Per concloure:
- Allunyar-nos del hype de les eines i acceptem que aquestes romandran d'una o altra manera (inferència local)
- Definir processos
- Capacitat de jutjar abans que velocitat, bones pràciques abans que eines, amplificar abans que acceptar i tenir responsabilitat plena
- L'enginyeria de programari és més viva que mai, tant per qui dissenya/teixeix sistemes com per qui en garanteix el seu bon funcionament
-->

---

<!-- _paginate: false -->

# De filadors de codi<br>a teixidors de sistemes

La pregunta no és si les màquines filaran per nosaltres.

La pregunta és **qui teixeix**.

<!--
[COMIAT — 1 minut]

- La revolució industrial tèxtil a Catalunya va ser molt més que un canvi tecnològic
- Va ser una fractura
- Les màquines van crear riquesa sense precedents, però el cost humà va ser immens
- Ara nosaltres vivim la revolució industrial del programari
- El codi ja es fila sol
- Per tant la pregunta és qui s'adaptarà i com
- I sobretot, qui s'endú el benefici de tot plegat
- Ens ho deien fa 200 anys des dels carrers de Barcelona: Associació o mort
- Jo ho traduiria per al 2026: Criteri o irrellevància
- Gràcies
-->

---

<!-- _paginate: false -->

<p class="small">

Dario Castañé — @darccio

cc@dario.cat
dario.cat

</p>

<!--
Diapositiva final amb contacte. No cal dir res, només deixar-la visible durant les preguntes
-->
