# TLS-Spoofer

## Что будем делать?

Делаем http-спуфер для обхода блокировок Cloud-Flare.  
Для обхода блокировки по TLS фингерпринту требуется подменить JA3 в http клиенте.

### Go

- Все кладем в папку `./go`
- Делаем **ws** *(предпочтительней)* или **http** сервер
- В качестве параметров запроса принимаем такую структуру:
  ```
  {
    *url: string
    *method: 'POST' | 'GET' | ...
    body: string
    cookies: {
      [key: string]: string
    }
    headers: {
      [key: string]: string
    },
    headersOrder: string[]
    ja3: string
    userAgent: string
    proxy: 'http://login:topsecret@192.168.1.1:80'
  }
  ```

  `* - обязательные параметры`
- На гошке создаем клиент с параметрами `ja3` и `userAgent`
- А после отправляем запрос на `url` согласно остальным параметрам запроса
- `headersOrder` указывает в каком порядке нужно поствавить заголовки
- `proxy` указывает через какой http прокси сервер нужно отрпавить запрос
- Ответ запроса нужно вернуть туда откуда пришел запрос (через `ws` или `http`) формат не принципиален
- Воркеров должно быть много, желательно это уметь настраивать через окружение
- При возникновении ошибки в воркере, он не доджен падать

### Проверить обход можно так

Если отправить GET запрос на такой урл `https://pathfinder.1inch.io/v1.4/chain/56/router/v5/quotes?fromTokenAddress=0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee&toTokenAddress=0x111111111117dc0aa78b770fa6a738034120c302&amount=553000000000000&gasPrice=5000000000&protocolWhiteList=PANCAKESWAP,VENUS,JULSWAP,BAKERYSWAP,BSC_ONE_INCH_LP,ACRYPTOS,BSC_DODO,APESWAP,SPARTAN_V2,VSWAP,VPEGSWAP,HYPERSWAP,BSC_DODO_V2,SWAPSWIPE,ELLIPSIS_FINANCE,BSC_NERVE,BSC_SMOOTHY_FINANCE,CHEESESWAP,BSC_PMM1,PANCAKESWAP_V2,MDEX,WARDEN,WAULTSWAP,BSC_ONE_INCH_LIMIT_ORDER,BSC_ONE_INCH_LIMIT_ORDER_V2,BSC_ONE_INCH_LIMIT_ORDER_V3,BSC_PMM3,BSC_PMM7,ACSI_FINANCE,GAMBIT_FINANCE,JETSWAP,BSC_UNIFI,BSC_PMMX,BSC_KYBER_DMM,BSC_BI_SWAP,BSC_DOPPLE,BABYSWAP,BSC_PMM2MM,WOOFI,BSC_ELK,BSC_SYNAPSE,BSC_AUTOSHARK,BSC_CAFE_SWAP,BSC_PMM5,PLANET_FINANCE,BSC_ANNEX_FINANCE,BSC_ANNEX_SWAP,BSC_RADIOSHACK,BSC_KYBERSWAP_ELASTIC,BSC_FSTSWAP,BSC_NOMISWAP,BSC_CONE,BSC_KYBER_DMM_STATIC,WOMBATSWAP,BSC_NOMISWAP_STABLE,BSC_PANCAKESWAP_STABLE,BSC_BABYDOGE,BSC_THENA,BSC_WOOFI_V2,BSC_KYOTOSWAP,BSC_TRADERJOE,BSC_TRADERJOE_V2,BSC_UNISWAP_V3,BSC_TRIDENT,BSC_MAVERICK_V1,BSC_PANCAKESWAP_V3,BSC_THENA_V3,BSC_PMM8,BSC_TRADERJOE_V2_1,BSC_NOMISWAPEPCS,BSC_USDFI,BSC_PMM11&walletAddress=0xb5af4d8251dbd1ae7623ae97a988da8f25cde124&preset=maxReturnResult
`

C заголовками:
```
headers: {
  'origin': 'https://app.1inch.io',
  'referer': 'https://app.1inch.io/'
  'accept': '*/*',
  'accept-encoding': 'gzip, deflate, br',
  'accept-language': 'ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7',
  'sec-ch-ua': '"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"',
  'connection': 'keep-alive',
  'sec-ch-ua-mobile': '?0',
  'sec-ch-ua-platform': '"macOS"',
  'sec-fetch-dest': 'empty',
  'sec-fetch-mode': 'cors',
  'sec-fetch-site': 'same-site',
}
```

JA3 таким: `771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,65281-10-51-35-16-43-17513-23-18-45-13-5-27-11-0-21,29-23-24,0`

И userAgent таким: `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 YaBrowser/23.5.4.696 Yowser/2.5 Safari/537.36`

А порядком заголовком таким
```
[
  'host',
  'connection',
  'cache-control',
  'device-memory',
  'viewport-width',
  'rtt',
  'downlink',
  'ect',
  'sec-ch-ua',
  'sec-ch-ua-mobile',
  'sec-ch-ua-full-version',
  'sec-ch-ua-arch',
  'sec-ch-ua-platform',
  'sec-ch-ua-platform-version',
  'sec-ch-ua-model',
  'upgrade-insecure-requests',
  'user-agent',
  'accept',
  'sec-fetch-site',
  'sec-fetch-mode',
  'sec-fetch-user',
  'sec-fetch-dest',
  'referer',
  'accept-encoding',
  'accept-language',
  'cookie',
]
```

То Cloud Flare должен пропустить такой запрос и вернуть ответ

### NodeJS

Здесь захуярим обвязочку для публикации пакета в npm и использовании в ноде. (Позже распишу)
