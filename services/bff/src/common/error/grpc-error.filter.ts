import {
  ArgumentsHost,
  BadRequestException,
  Catch,
  ExceptionFilter,
  InternalServerErrorException,
  Logger,
  NotFoundException,
  UnauthorizedException,
} from '@nestjs/common'
import { GqlArgumentsHost } from '@nestjs/graphql'
import { isGrpcError } from './grpc-error.type'
import { status } from '@grpc/grpc-js'
import { GraphQLResolveInfo } from 'graphql'

@Catch()
export class GrpcExceptionFilter implements ExceptionFilter {
  private readonly logger = new Logger(GrpcExceptionFilter.name, {
    timestamp: true,
  })

  catch(exception: unknown, host: ArgumentsHost) {
    const gqlHost = GqlArgumentsHost.create(host)
    const info = gqlHost.getInfo<GraphQLResolveInfo>()
    const args = gqlHost.getArgs<Record<string, unknown>>()

    if (isGrpcError(exception)) {
      const message =
        `[GraphQL:${info.parentType.name}.${info.fieldName}] gRPC Error: ` +
        `code=${exception.code}, message="${exception.message}", ` +
        `args=${JSON.stringify(args)}`

      switch (exception.code) {
        case status.INVALID_ARGUMENT:
        case status.NOT_FOUND:
        case status.UNAUTHENTICATED:
          this.logger.log(message)
          break
        default:
          this.logger.error(message)
      }

      switch (exception.code) {
        case status.INVALID_ARGUMENT:
          throw new BadRequestException(exception.message)
        case status.NOT_FOUND:
          throw new NotFoundException(exception.message)
        case status.UNAUTHENTICATED:
          throw new UnauthorizedException(exception.message)
        default:
          throw new InternalServerErrorException(exception.message)
      }
    }

    throw exception
  }
}
