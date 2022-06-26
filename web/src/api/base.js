import request from '@/utils/request';

export const getConnList = query => {
    return request({
        url: 'api/conn/list',
        method: 'get',
        params: query
    });
};