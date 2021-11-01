import { createI18n } from "vue-i18n"
import zhCnLocale from 'element-plus/lib/locale/lang/zh-cn';

// 定义语言国际化内容
/**
 * 说明：
 * /src/i18n/languages 下的 ts 为框架的国际化内容
 */
export function loadLanguages() {
  const context = import.meta.globEager("./languages/*.ts");
  const languages: AnyObject = {};
  let langList = Object.keys(context);
  for (let key of langList) {
    if (key === "./index.ts") return;
    let lang = context[key].lang;
    let name = key.replace(/(\.\/languages\/|\.ts)/g, '');
    languages[name] = lang
  }
  return languages
}

export const i18n = createI18n({
  // globalInjection: true,
  // legacy: false,
  locale: zhCnLocale.name,
  fallbackLocale: zhCnLocale.name,
  messages: loadLanguages()
})

export function setLanguage(locale: string) {
  i18n.global.locale = locale
}
