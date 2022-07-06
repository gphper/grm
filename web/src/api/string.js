import request from '@/utils/request';

export const showString = data => {
    return request({
        url: 'api/string/show',
        method: 'post',
        data
    });
};