/* tslint:disable */
/* eslint-disable */
/**
 * internal nupin API
 * The specification for the internal API of nupin
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * An error response.
 * @export
 * @interface ApiError
 */
export interface ApiError {
    /**
     * 
     * @type {string}
     * @memberof ApiError
     */
    error: string;
    /**
     * 
     * @type {string}
     * @memberof ApiError
     */
    displayMessage?: string;
}

/**
 * Check if a given object implements the ApiError interface.
 */
export function instanceOfApiError(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "error" in value;

    return isInstance;
}

export function ApiErrorFromJSON(json: any): ApiError {
    return ApiErrorFromJSONTyped(json, false);
}

export function ApiErrorFromJSONTyped(json: any, ignoreDiscriminator: boolean): ApiError {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'error': json['error'],
        'displayMessage': !exists(json, 'displayMessage') ? undefined : json['displayMessage'],
    };
}

export function ApiErrorToJSON(value?: ApiError | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'error': value.error,
        'displayMessage': value.displayMessage,
    };
}

