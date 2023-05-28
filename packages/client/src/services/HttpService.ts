export interface FetchResult<T> extends Response {
  data?: T;
}

export type CustomRequestInit = Omit<RequestInit, "method" | "body"> & {
  body?: unknown;
};

export default class HttpService {
  private static transformRequestInit(
    init: CustomRequestInit,
    method?: string
  ): RequestInit {
    const token = localStorage.getItem("accessToken");
    const Authorization = `Bearer ${token}`;
    return {
      ...init,
      method,
      body: JSON.stringify(init.body),
      credentials: "include",
      headers: {
        ...init.headers,
        "Content-Type": "application/json",
        Authorization,
      },
    };
  }

  private static async request<T>(
    url: string,
    init: CustomRequestInit = {},
    method?: string
  ): Promise<FetchResult<T>> {
    const transformedInit = this.transformRequestInit(init, method);
    const res = await fetch(url, transformedInit);
    const bodyLength = Number(res.headers.get("Content-Length"));
    const resWithData: FetchResult<T> = res;
    try {
      const data: T = bodyLength > 0 ? await res.json() : undefined;
      resWithData.data = data;
    } catch (e) {
      console.error("Failed to get json from request: " + e);
    }

    return resWithData;
  }

  public static async get<T>(
    url: string,
    init?: CustomRequestInit
  ): Promise<FetchResult<T>> {
    return this.request<T>(url, init, "GET");
  }

  public static async post<T>(
    url: string,
    init?: CustomRequestInit
  ): Promise<FetchResult<T>> {
    return this.request<T>(url, init, "POST");
  }

  public static async put<T>(
    url: string,
    init?: CustomRequestInit
  ): Promise<FetchResult<T>> {
    return this.request<T>(url, init, "PUT");
  }

  public static async patch<T>(
    url: string,
    init?: CustomRequestInit
  ): Promise<FetchResult<T>> {
    return this.request<T>(url, init, "PATCH");
  }

  public static async delete<T>(
    url: string,
    init?: CustomRequestInit
  ): Promise<FetchResult<T>> {
    return this.request<T>(url, init, "DELETE");
  }
}
