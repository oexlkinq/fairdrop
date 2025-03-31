/* tslint:disable */
/* eslint-disable */
/**
 * fairdrop api
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from './configuration';
import type { AxiosPromise, AxiosInstance, RawAxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
import type { RequestArgs } from './base';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, BaseAPI, RequiredError, operationServerMap } from './base';

/**
 * 
 * @export
 * @interface FoldersCreatePost200Response
 */
export interface FoldersCreatePost200Response {
    /**
     * 
     * @type {string}
     * @memberof FoldersCreatePost200Response
     */
    'password': string;
    /**
     * timestamp в формате unix-time в секундах
     * @type {number}
     * @memberof FoldersCreatePost200Response
     */
    'creation_date': number;
    /**
     * 
     * @type {string}
     * @memberof FoldersCreatePost200Response
     */
    'ip': string;
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * создаёт новую папку и возвращает информацию о ней
         * @summary создать папку
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        foldersCreatePost: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/folders/create`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * удаляет папку
         * @summary удалить папку
         * @param {string} folder пароль-идентификатор папки
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        foldersFolderDelete: async (folder: string, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'folder' is not null or undefined
            assertParamExists('foldersFolderDelete', 'folder', folder)
            const localVarPath = `/folders/{folder}`
                .replace(`{${"folder"}}`, encodeURIComponent(String(folder)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'DELETE', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary скачать файл
         * @param {string} folder пароль-идентификатор папки
         * @param {string} filename имя файла
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        foldersFolderFilenameGet: async (folder: string, filename: string, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'folder' is not null or undefined
            assertParamExists('foldersFolderFilenameGet', 'folder', folder)
            // verify required parameter 'filename' is not null or undefined
            assertParamExists('foldersFolderFilenameGet', 'filename', filename)
            const localVarPath = `/folders/{folder}/{filename}`
                .replace(`{${"folder"}}`, encodeURIComponent(String(folder)))
                .replace(`{${"filename"}}`, encodeURIComponent(String(filename)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * возвращает список из имен файлов содержащихся в папке
         * @summary получить список файлов папки
         * @param {string} folder пароль-идентификатор папки
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        foldersFolderGet: async (folder: string, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'folder' is not null or undefined
            assertParamExists('foldersFolderGet', 'folder', folder)
            const localVarPath = `/folders/{folder}`
                .replace(`{${"folder"}}`, encodeURIComponent(String(folder)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * записывает файлы переданные в форме в указанную папку
         * @summary загрузить файлы в папку
         * @param {string} folder пароль-идентификатор папки
         * @param {Array<File>} [file] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        foldersFolderPost: async (folder: string, file?: Array<File>, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'folder' is not null or undefined
            assertParamExists('foldersFolderPost', 'folder', folder)
            const localVarPath = `/folders/{folder}`
                .replace(`{${"folder"}}`, encodeURIComponent(String(folder)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;
            const localVarFormParams = new ((configuration && configuration.formDataCtor) || FormData)();

            if (file) {
                file.forEach((element) => {
                    localVarFormParams.append('file', element as any);
                })
            }

    
    
            localVarHeaderParameter['Content-Type'] = 'multipart/form-data';
    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = localVarFormParams;

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * DefaultApi - functional programming interface
 * @export
 */
export const DefaultApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = DefaultApiAxiosParamCreator(configuration)
    return {
        /**
         * создаёт новую папку и возвращает информацию о ней
         * @summary создать папку
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async foldersCreatePost(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<FoldersCreatePost200Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.foldersCreatePost(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.foldersCreatePost']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * удаляет папку
         * @summary удалить папку
         * @param {string} folder пароль-идентификатор папки
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async foldersFolderDelete(folder: string, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.foldersFolderDelete(folder, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.foldersFolderDelete']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary скачать файл
         * @param {string} folder пароль-идентификатор папки
         * @param {string} filename имя файла
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async foldersFolderFilenameGet(folder: string, filename: string, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<File>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.foldersFolderFilenameGet(folder, filename, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.foldersFolderFilenameGet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * возвращает список из имен файлов содержащихся в папке
         * @summary получить список файлов папки
         * @param {string} folder пароль-идентификатор папки
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async foldersFolderGet(folder: string, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<string>>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.foldersFolderGet(folder, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.foldersFolderGet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * записывает файлы переданные в форме в указанную папку
         * @summary загрузить файлы в папку
         * @param {string} folder пароль-идентификатор папки
         * @param {Array<File>} [file] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async foldersFolderPost(folder: string, file?: Array<File>, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.foldersFolderPost(folder, file, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.foldersFolderPost']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
    }
};

/**
 * DefaultApi - factory interface
 * @export
 */
export const DefaultApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = DefaultApiFp(configuration)
    return {
        /**
         * создаёт новую папку и возвращает информацию о ней
         * @summary создать папку
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        foldersCreatePost(options?: RawAxiosRequestConfig): AxiosPromise<FoldersCreatePost200Response> {
            return localVarFp.foldersCreatePost(options).then((request) => request(axios, basePath));
        },
        /**
         * удаляет папку
         * @summary удалить папку
         * @param {string} folder пароль-идентификатор папки
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        foldersFolderDelete(folder: string, options?: RawAxiosRequestConfig): AxiosPromise<void> {
            return localVarFp.foldersFolderDelete(folder, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary скачать файл
         * @param {string} folder пароль-идентификатор папки
         * @param {string} filename имя файла
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        foldersFolderFilenameGet(folder: string, filename: string, options?: RawAxiosRequestConfig): AxiosPromise<File> {
            return localVarFp.foldersFolderFilenameGet(folder, filename, options).then((request) => request(axios, basePath));
        },
        /**
         * возвращает список из имен файлов содержащихся в папке
         * @summary получить список файлов папки
         * @param {string} folder пароль-идентификатор папки
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        foldersFolderGet(folder: string, options?: RawAxiosRequestConfig): AxiosPromise<Array<string>> {
            return localVarFp.foldersFolderGet(folder, options).then((request) => request(axios, basePath));
        },
        /**
         * записывает файлы переданные в форме в указанную папку
         * @summary загрузить файлы в папку
         * @param {string} folder пароль-идентификатор папки
         * @param {Array<File>} [file] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        foldersFolderPost(folder: string, file?: Array<File>, options?: RawAxiosRequestConfig): AxiosPromise<void> {
            return localVarFp.foldersFolderPost(folder, file, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * DefaultApi - object-oriented interface
 * @export
 * @class DefaultApi
 * @extends {BaseAPI}
 */
export class DefaultApi extends BaseAPI {
    /**
     * создаёт новую папку и возвращает информацию о ней
     * @summary создать папку
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public foldersCreatePost(options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).foldersCreatePost(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * удаляет папку
     * @summary удалить папку
     * @param {string} folder пароль-идентификатор папки
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public foldersFolderDelete(folder: string, options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).foldersFolderDelete(folder, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary скачать файл
     * @param {string} folder пароль-идентификатор папки
     * @param {string} filename имя файла
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public foldersFolderFilenameGet(folder: string, filename: string, options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).foldersFolderFilenameGet(folder, filename, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * возвращает список из имен файлов содержащихся в папке
     * @summary получить список файлов папки
     * @param {string} folder пароль-идентификатор папки
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public foldersFolderGet(folder: string, options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).foldersFolderGet(folder, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * записывает файлы переданные в форме в указанную папку
     * @summary загрузить файлы в папку
     * @param {string} folder пароль-идентификатор папки
     * @param {Array<File>} [file] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public foldersFolderPost(folder: string, file?: Array<File>, options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).foldersFolderPost(folder, file, options).then((request) => request(this.axios, this.basePath));
    }
}



