.global _start

_start:
	  mov x0, #1
	  adr x1, hello_world
	  mov x2, #14
	  mov x8, #64
	  svc 0

	  mov x8, #93
	  mov x0, #0
	  svc 0

hello_world:
	  .asciz "Hello, World!\n"
