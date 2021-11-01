import request from '/@/utils/request'
import { AxiosRequestConfig } from "axios";

export const getRequest = (url: string, params?: object) => {
  return request({
    method: 'get',
    url: url,
    params: params
  })
}

export const getRequestById = (url: string, config?: AxiosRequestConfig) => {
  return request({
    method: 'get',
    url: url,
    ...config
  })
}

export const postRequest = (url: string, data?: object) => {
  return request({
    method: 'post',
    url: url,
    data: data
  })
}

export const putRequest = (url: string, data?: object) => {
  return request({
    method: 'put',
    url: url,
    data: data
  })
}

export const deleteRequest = (url: string, data?: object) => {
  return request({
    method: 'delete',
    url: url,
    data: data
  })
}

export const deleteRequestById = (url: string) => {
  return request({
    method: 'delete',
    url: url
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
