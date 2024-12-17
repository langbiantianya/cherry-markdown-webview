

export default {
  "meta": {},
  "id": "_default",
  "name": "",
  "file": {
    "path": "src/routes",
    "dir": "src",
    "base": "routes",
    "ext": "",
    "name": "routes"
  },
  "rootName": "default",
  "routifyDir": import.meta.url,
  "children": [
    {
      "meta": {
        "isDefault": true
      },
      "id": "_default_index_svelte",
      "name": "index",
      "file": {
        "path": "src/routes/index.svelte",
        "dir": "src/routes",
        "base": "index.svelte",
        "ext": ".svelte",
        "name": "index"
      },
      "asyncModule": () => import('../src/routes/index.svelte'),
      "children": []
    },
    {
      "meta": {},
      "id": "_default_settings",
      "name": "settings",
      "module": false,
      "file": {
        "path": "src/routes/settings",
        "dir": "src/routes",
        "base": "settings",
        "ext": "",
        "name": "settings"
      },
      "children": [
        {
          "meta": {},
          "id": "_default_settings_options",
          "name": "options",
          "module": false,
          "file": {
            "path": "src/routes/settings/options",
            "dir": "src/routes/settings",
            "base": "options",
            "ext": "",
            "name": "options"
          },
          "children": [
            {
              "meta": {
                "isDefault": true
              },
              "id": "_default_settings_options_index_svelte",
              "name": "index",
              "file": {
                "path": "src/routes/settings/options/index.svelte",
                "dir": "src/routes/settings/options",
                "base": "index.svelte",
                "ext": ".svelte",
                "name": "index"
              },
              "asyncModule": () => import('../src/routes/settings/options/index.svelte'),
              "children": []
            }
          ]
        },
        {
          "meta": {},
          "id": "_default_settings_personaliza",
          "name": "personaliza",
          "module": false,
          "file": {
            "path": "src/routes/settings/personaliza",
            "dir": "src/routes/settings",
            "base": "personaliza",
            "ext": "",
            "name": "personaliza"
          },
          "children": [
            {
              "meta": {
                "isDefault": true
              },
              "id": "_default_settings_personaliza_index_svelte",
              "name": "index",
              "file": {
                "path": "src/routes/settings/personaliza/index.svelte",
                "dir": "src/routes/settings/personaliza",
                "base": "index.svelte",
                "ext": ".svelte",
                "name": "index"
              },
              "asyncModule": () => import('../src/routes/settings/personaliza/index.svelte'),
              "children": []
            }
          ]
        }
      ]
    },
    {
      "meta": {
        "dynamic": true,
        "dynamicSpread": true,
        "order": false,
        "inline": false
      },
      "name": "[...404]",
      "file": {
        "path": ".routify/components/[...404].svelte",
        "dir": ".routify/components",
        "base": "[...404].svelte",
        "ext": ".svelte",
        "name": "[...404]"
      },
      "asyncModule": () => import('./components/[...404].svelte'),
      "children": []
    }
  ]
}