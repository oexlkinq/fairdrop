import { Configuration, DefaultApi } from "./client";

const publicBasePath = import.meta.env.VITE_PUBLIC_BASE_PATH as string | undefined

const conf = new Configuration({
  basePath: publicBasePath ?? import.meta.env.BASE_URL
})
export const api = new DefaultApi(conf)
export default api
