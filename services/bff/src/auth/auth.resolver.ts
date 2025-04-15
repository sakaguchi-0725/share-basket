import { Args, Mutation, Resolver, Context } from '@nestjs/graphql'
import { AuthService } from './auth.service'
import { LoginInput, LoginOutput } from './dto/login.dto'
import { SignUpInput, SignUpOutput } from './dto/signup.dto'
import {
  SignUpConfirmInput,
  SignUpConfirmOutput,
} from './dto/signup-confirm.dto'
import { Response } from 'express'

@Resolver()
export class AuthResolver {
  constructor(private readonly authService: AuthService) {}

  @Mutation(() => LoginOutput)
  async login(
    @Args('input') input: LoginInput,
    @Context() context: { res: Response },
  ): Promise<LoginOutput> {
    console.log('login')
    const token = await this.authService.login(input)

    context.res.cookie('accessToken', token, {
      httpOnly: true,
      secure: process.env.NODE_ENV === 'production', // 本番環境ではセキュアフラグを設定
      sameSite: 'strict',
      maxAge: 24 * 60 * 60 * 1000,
    })

    return {
      message: 'ok',
    }
  }

  @Mutation(() => SignUpOutput)
  signup(@Args('input') input: SignUpInput): Promise<SignUpOutput> {
    return this.authService.signup(input)
  }

  @Mutation(() => SignUpConfirmOutput)
  signUpConfirm(
    @Args('input') input: SignUpConfirmInput,
  ): Promise<SignUpConfirmOutput> {
    return this.authService.signupConfirm(input)
  }
}
