import { Field, InputType, ObjectType } from '@nestjs/graphql'

@InputType()
export class SignUpConfirmInput {
  @Field()
  email: string

  @Field()
  confirmationCode: string
}

@ObjectType()
export class SignUpConfirmOutput {
  @Field()
  message: string
}
