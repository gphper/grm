import request from '@/utils/request';

export const openDb = data => {
    return request({
        url: 'api/index/open',
        method: 'post',
        data
    });
};