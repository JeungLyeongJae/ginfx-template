export class AuthConfigCosts {
    public static DEFAULT_TOKEN_NAME = 'token';
    public static DEFAULT_HEADER_NAME = 'Authorization';
    public static HEADER_PREFIX_BEARER = 'Bearer ';
}


export class JwtHelper {

    public base64UrlDecode(base64Url: string): string {
        // Replace non-url compatible chars with base64 standard chars
        base64Url = base64Url.replace(/-/g, '+').replace(/_/g, '/');

        // Pad out with standard base64 required padding characters
        switch (base64Url.length % 4) {
            case 0:
                break;
            case 2:
                base64Url += '==';
                break;
            case 3:
                base64Url += '=';
                break;
            default:
                throw new Error('Invalid base64url string!');
        }

        // Decode base64 string
        return atob(base64Url);
    }

    public decodeToken(token: string): any {
        const parts = token.split('.');

        if (parts.length !== 3) {
            throw new Error('JWT must have 3 parts');
        }

        const decoded = this.base64UrlDecode(parts[1]);
        if (!decoded) {
            throw new Error('Cannot decode the token');
        }

        return JSON.parse(decoded);
    }

    public getTokenExpirationDate(token: string): Date | null {
        let decoded: any;
        decoded = this.decodeToken(token);

        if (!decoded.hasOwnProperty('exp')) {
            return null;
        }

        const date = new Date(0); // The 0 here is the key, which sets the date to the epoch
        date.setUTCSeconds(decoded.exp);

        return date;
    }

    public isTokenExpired(token: string, offsetSeconds?: number): boolean {
        const date = this.getTokenExpirationDate(token);
        offsetSeconds = offsetSeconds || 0;

        if (date == null) {
            return false;
        }

        // Token expired?
        return !(date.valueOf() > (new Date().valueOf() + (offsetSeconds * 1000)));
    }
}

/**
 * Checks for presence of token and that token hasn't expired.
 * For use with the @CanActivate router decorator and NgIf
 */
export function tokenNotExpired(tokenName = AuthConfigCosts.DEFAULT_TOKEN_NAME, jwt?: string): boolean {

    const token: string | null = jwt || sessionStorage.getItem(tokenName);

    const jwtHelper = new JwtHelper();

    return token != null && !jwtHelper.isTokenExpired(token);
}