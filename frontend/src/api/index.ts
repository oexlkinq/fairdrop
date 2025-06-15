import { Configuration, DefaultApi } from "./client";

const publicBasePath = import.meta.env.VITE_PUBLIC_BASE_PATH as string | undefined

const conf = new Configuration({
  basePath: publicBasePath
})
export const api = new DefaultApi(conf)
export default api
