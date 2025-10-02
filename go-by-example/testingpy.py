import pygame
pygame.init()

scrn = pygame.display.set_mode((400, 300))
pygame.display.set_caption("Test window")

running = True
while running:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False


pygame.quit()
