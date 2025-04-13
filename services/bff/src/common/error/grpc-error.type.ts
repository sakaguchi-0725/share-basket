import { status } from '@grpc/grpc-js'

export type GrpcError = {
  code: (typeof status)[keyof typeof status]
  message: string
}

export const isGrpcError = (err: unknown): err is GrpcError => {
  return (
    typeof err === 'object' &&
    err !== null &&
    'code' in err &&
    typeof err.code === 'number' &&
    'message' in err &&
    typeof err.message === 'string'
  )
}
