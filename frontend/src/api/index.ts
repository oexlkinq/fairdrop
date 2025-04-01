import { Configuration, DefaultApi } from "./client";

const conf = new Configuration({
  // #2 TODO: этот адрес стоит доставать из энва например
  basePath: 'http://localhost:8081'
})
export const api = new DefaultApi(conf)
export default api
