import { Args, Mutation, Query, Resolver } from '@nestjs/graphql'
import { UseGuards } from '@nestjs/common'
import { ShoppingItem } from './dto/shopping-item.dto'
import { CreateShoppingItemInput } from './dto/create.input'
import { PersonalShoppingService } from './personal-shopping.service'
import { AuthGuard } from '../auth/auth.guard'

@Resolver()
@UseGuards(AuthGuard)
export class PersonalShoppingResolver {
  constructor(private readonly shoppingService: PersonalShoppingService) {}

  @Query(() => [ShoppingItem])
  getAll(): Promise<ShoppingItem[]> {
    return this.shoppingService.getAll()
  }

  @Mutation(() => ShoppingItem)
  create(@Args('input') input: CreateShoppingItemInput): Promise<ShoppingItem> {
    return this.shoppingService.create(input)
  }
}
