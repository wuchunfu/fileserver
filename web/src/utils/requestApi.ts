import request from '/@/utils/request'
import { AxiosRequestConfig } from "axios";

export const getRequest = (url: string, params?: object, config?: AxiosRequestConfig) => {
  return request({
    method: 'get',
    url: url,
    params: params,
    ...config
  })
}

export const getRequestById = (url: string, config?: AxiosRequestConfig) => {
  return request({
    method: 'get',
    url: url,
    ...config
  })
}

export const postRequest = (url: string, data?: object, config?: AxiosRequestConfig) => {
  return request({
    method: 'post',
    url: url,
    data: data,
    ...config
  })
}

export const putRequest = (url: string, data?: object, config?: AxiosRequestConfig) => {
  return request({
    method: 'put',
    url: url,
    data: data,
    ...config
  })
}

export const deleteRequest = (url: string, data?: object, config?: AxiosRequestConfig) => {
  return request({
    method: 'delete',
    url: url,
    data: data,
    ...config
  })
}

export const deleteRequestById = (url: string, config?: AxiosRequestConfig) => {
  return request({
    method: 'delete',
    url: url,
    ...config
  })
}

export const uploadFileRequest = (url: string, data?: object, config?: AxiosRequestConfig) => {
  return request({
    method: 'post',
    url: url,
    data: data,
    ...config
  })
}
