import {
  CanActivate,
  ExecutionContext,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common'
import { Observable } from 'rxjs'
import { GqlExecutionContext } from '@nestjs/graphql'
import { Request } from 'express'
import { AuthService } from './auth.service'

interface GqlContext {
  req: Request
}

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(private readonly authService: AuthService) {}

  canActivate(
    context: ExecutionContext,
  ): boolean | Promise<boolean> | Observable<boolean> {
    const gqlContext =
      GqlExecutionContext.create(context).getContext<GqlContext>()
    const request = gqlContext.req

    if (!request?.cookies?.accessToken) {
      throw new UnauthorizedException('アクセストークンが見つかりません')
    }

    // トークンを検証
    const token = request.cookies.accessToken as string
    return this.verifyToken(token)
  }

  private verifyToken(token: string): boolean {
    // TODO: トークン検証を行う
    console.log(token)
    return true
  }
}
