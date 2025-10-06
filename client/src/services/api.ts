const API_BASE_URL = 'http://localhost:3333'

export class ApiService {
  private static async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    const url = `${API_BASE_URL}${endpoint}`
    const config: RequestInit = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
      ...options,
    }

    try {
      const response = await fetch(url, config)
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      
      return await response.json()
    } catch (error) {
      console.error('API request failed:', error)
      throw error
    }
  }

  static async getAuthUrl(): Promise<{ url: string }> {
    return this.request<{ url: string }>('/auth/url', {
      method: 'GET',
    })
  }

  static async verifyToken(token: string): Promise<{ valid: boolean; email?: string }> {
    return this.request<{ valid: boolean; email?: string }>('/auth/verify', {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    })
  }
}
